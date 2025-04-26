// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {PlonkVerifier} from "./PlonkVerifier.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

/**
 * @title FoundnoneVRF
 * @author Zachary Owens (foundnone.eth)
 * @notice A democratized VRF that allows anyone to request entropy as well as fulfill entropy requests for a reward.
 * @notice This contract will refer to the msg.sender as the fulfiller and the address that will receive the reward as the rewardReceiver.
 * @dev Entropy emitted by this contract is verified by the PlonkVerifier before being stored in the contract.
 * @dev This allows the reward receiver to be different from the fulfiller.
 * @dev To avoid spoofing, the contract requires a predetermined commitment to be passed in by the fulfiller.
 * @dev This commitment is used as a public input to the PlonkVerifier and is used to verify the proof.
 * @dev The commitment is the hash of a secret that the fulfiller must use when generating the next entropy.
 * @dev Additionally, the seed of the request must always match the keccak256(requestId, msg.sender) hash.
 * @dev Since all public inputs are constrained, and the pre-hash commitment is predetermined, we can trust that the entropy (if valid) is properly generated and not tampered with.
 */

contract FoundnoneVRF is PlonkVerifier, AccessControl {
    /**
     * @notice The admin role for the contract
     * @dev The admin role is used to manage the request fee and contract fee percentage and to withdraw contract fees
     */
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /**
     * @notice A mapping to track if a request has been fulfilled and to easily retrieve the entropy
     */
    mapping(uint256 => uint256) public entropies;

    mapping(uint256 => uint256) public requestBlockSet;

    /**
     * @notice The next requestId to be used
     * @dev This is incremented every time a new request is made
     */
    uint256 public nextRequestId = 1;

    /**
     * @notice The fee required to make a request
     */
    uint256 public requestFee = 0.000005 ether;

    /**
     * @notice The percentage of the fee that goes to the contract owner
     */
    uint256 public contractFeePercentage = 5;

    /**
     * @notice A mapping to track the contract balance of each reward receiver
     */
    mapping(address => uint256) public rewardReceiverBalance;

    /**
     * @notice A mapping to store the commitments for each fulfiller
     * @dev A commitment is a hash of a secret that the fulfiller must use when generating the next entropy
     * @dev When fulfilling a request, a new commitment should be passed to the contract to set the new commitment for that fulfiller
     * @dev This is used to ensure that fulfillers are not able to spoof entropies
     */
    mapping(address => uint256) public commitments;

    /**
     * @notice A mapping to track the block number when a commitment was set
     * @dev This is used to ensure that fulfillers are only using commitments on requests that are made after the commitment is set
     */
    mapping(address => uint256) public commitmentBlockSet;

    /**
     * @notice A mapping to track the contract fee balance
     */
    uint256 public contractFeeBalance;

    /**
     * @notice Event emitted when a request is fulfilled, the fulfilling address is the fee receiver
     * @param requestId The requestId of the request
     * @param rewardReceiver The address to receive the reward
     * @param proof The proof data
     * @param publicInputs The public inputs of the proof
     */
    event EntropyStored(
        uint256 indexed requestId,
        address rewardReceiver,
        uint256[24] proof,
        uint256[3] publicInputs
    );

    /**
     * @notice Event emitted when a request is made
     * @param requestId The requestId of the request
     * @param requester The address of the requester
     * @param feePaid The fee paid for the request
     */
    event VrfRequested(uint256 requestId, address requester, uint256 feePaid);

    /**
     * @notice Event emitted when the contract fee balance is withdrawn
     * @param amount The amount withdrawn
     */
    event ContractFeesWithdrawn(uint256 amount);

    /**
     * @notice Event emitted when a reward receiver balance is withdrawn
     * @param rewardReceiver The address of the reward receiver
     * @param amount The amount withdrawn
     */
    event RewardReceiverBalanceWithdrawn(address indexed rewardReceiver, uint256 amount);

    /**
     * @notice Event emitted when the request fee is updated
     * @param newFee The new fee
     */
    event RequestFeeUpdated(uint256 newFee);

    /**
     * @notice Event emitted when the contract fee percentage is updated
     * @param newPercentage The new percentage
     */
    event ContractFeePercentageUpdated(uint256 newPercentage);

    error InvalidRequestId();
    error RequestAlreadyFulfilled();
    error InvalidProof();
    error InvalidCommitment();
    error CommitmentAlreadySet();
    error InvalidSeed();
    error InvalidCommitmentBlock();
    error DuplicateCommitment();
    error InsufficientFee();
    error RequestNotFulfilled();
    error InvalidPercentage();
    error InsufficientBalance();

    /**
     * @notice The constructor for the contract
     * @param _adminRole The address of the admin role
     */
    constructor(address _adminRole) {
        _grantRole(DEFAULT_ADMIN_ROLE, _adminRole);
        _grantRole(ADMIN_ROLE, _adminRole);
    }

    /////////////////// Fulfiller functions ///////////////////////
    /**
     * @notice A function to verify the proof and store the entropy
     * @param _proof The proof data
     * @param _publicInputs The public inputs of the proof
     * @param _requestId The requestId of the request
     * @param _rewardReceiver The address to receive the reward
     * @param _nextCommitment The next commitment of the fulfiller
     */
    function submitEntropy(
        uint256[24] calldata _proof,
        uint256[3] calldata _publicInputs,
        uint256 _requestId,
        address _rewardReceiver,
        uint256 _nextCommitment
    ) external {
        if (commitments[msg.sender] == _nextCommitment) {
            revert DuplicateCommitment();
        }
        // verify the proof
        _verifyProof(_proof, _publicInputs, _requestId);

        // store the entropy
        entropies[_requestId] = _publicInputs[1];
    
        commitments[msg.sender] = _nextCommitment;
        commitmentBlockSet[msg.sender] = block.number;

        // calculate the fee
        uint256 fee = (requestFee * contractFeePercentage) / 100;

        // update balances
        contractFeeBalance += fee;
        rewardReceiverBalance[_rewardReceiver] += requestFee - fee;

        // emit the event
        emit EntropyStored(_requestId, _rewardReceiver, _proof, _publicInputs);
    }

    /**
     * @notice An all in one verifier function
     * @dev First, it checks if the requestId is valid and if the request has not been fulfilled yet
     * @dev Then, it checks if the commitment is valid and if the seed is valid
     * @dev Finally, it checks if the proof is valid using the PlonkVerifier which in turn verifies that the secret used in combination
     * with the public inputs is valid
     * @param _proof the proof data
     * @param _publicInputs the public inputs of the proof
     * @param _requestId the requestId of the request
     */
    function _verifyProof(
        uint256[24] calldata _proof,
        uint256[3] calldata _publicInputs,
        uint256 _requestId
    ) internal view {
        if (_requestId > nextRequestId) {
            revert InvalidRequestId();
        }
        if (entropies[_requestId] != 0) {
            revert RequestAlreadyFulfilled();
        }
        if (commitments[msg.sender] != (_publicInputs[2])) {
            revert InvalidCommitment();
        }
        if (commitmentBlockSet[msg.sender] > requestBlockSet[_requestId]) {
            revert InvalidCommitmentBlock();
        }
        uint256 _seedHash = uint256(
            keccak256(abi.encodePacked(_requestId, requestBlockSet[_requestId]))
        );
        // q is inherited from the PlonkVerifier contract and is the maximum value of the field
        uint256 _fieldHash = _seedHash % q;
        if (_publicInputs[0] != _fieldHash) {
            revert InvalidSeed();
        }
        if (!this.verifyProof(_proof, _publicInputs)) {
            revert InvalidProof();
        }
    }

    /**
     * @param _commitment The commitment to be set
     */
    function setCommitment(uint256 _commitment) external {
        if (commitments[msg.sender] == _commitment) {
            revert DuplicateCommitment();
        }
        commitments[msg.sender] = _commitment;
        commitmentBlockSet[msg.sender] = block.number;
    }

    /**
     * @notice A function to get the balance of a given reward receiver
     * @param _rewardReceiver The address to get the balance of
     * @return The balance of the given reward receiver
     */
    function getRewardReceiverBalance(
        address _rewardReceiver
    ) external view returns (uint256) {
        return rewardReceiverBalance[_rewardReceiver];
    }

    /**
     * @notice A function to withdraw the balance of a given reward receiver
     * @dev Must be called by the reward receiver that has a balance
     */
    function withdrawRewardReceiverBalance() external {
        uint256 balance = rewardReceiverBalance[msg.sender];
        if (balance == 0) {
            revert InsufficientBalance();
        }
        rewardReceiverBalance[msg.sender] = 0;
        payable(msg.sender).transfer(balance);
        emit RewardReceiverBalanceWithdrawn(msg.sender, balance);
    }

    /////////////////// Requester function ///////////////////////
    /**
     * @notice A function to request a new entropy
     * @dev The fee is paid in ether and is required to be sent with the request
     */
    function requestRng() external payable {
        if (msg.value < requestFee) {
            revert InsufficientFee();
        }
        requestBlockSet[nextRequestId] = block.number;
        emit VrfRequested(nextRequestId++, msg.sender, msg.value);
    }

    /**
     * @notice A function to get the entropy for a given requestId
     * @param _requestId The requestId of the request
     * @return The entropy generated from the seed
     */
    function getEntropy(uint256 _requestId) external view returns (uint256) {
        if (entropies[_requestId] == 0) {
            revert RequestNotFulfilled();
        }
        return entropies[_requestId];
    }

    /////////////////// Admin functions ///////////////////////
    /**
     * @param _newFee The new fee to be set
     * @dev This function can only be called by the admin role
     */
    function setRequestFee(uint256 _newFee) external onlyRole(ADMIN_ROLE) {
        requestFee = _newFee;
        emit RequestFeeUpdated(_newFee);
    }

    /**
     * @notice A function to withdraw the contract fee balance
     * @dev This function can only be called by the admin role
     */
    function withdraw() external onlyRole(ADMIN_ROLE) {
        payable(msg.sender).transfer(address(this).balance);
        contractFeeBalance = 0;
        emit ContractFeesWithdrawn(address(this).balance);
    }

    /**
     * @param _newPercentage The new percentage to be set
     * @dev This function can only be called by the admin role
     */
    function setContractFeePercentage(
        uint256 _newPercentage
    ) external onlyRole(ADMIN_ROLE) {
        if (_newPercentage > 20) {
            revert InvalidPercentage();
        }
        contractFeePercentage = _newPercentage;
        emit ContractFeePercentageUpdated(_newPercentage);
    }
}
