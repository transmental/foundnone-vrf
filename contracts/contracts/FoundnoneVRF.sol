// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {PlonkVerifier} from "./PlonkVerifier.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

import {IFoundnoneVRF} from "./interfaces/IFoundnoneVRF.sol";

/**
 * @title FoundnoneVRF
 * @author Zachary Owens (foundnone.eth)
 * @notice A democratized VRF that allows anyone to request entropy as well as fulfill entropy requests for a reward.
 * @notice This contract will refer to the msg.sender as the fulfiller and the address that will receive the reward as the rewardReceiver. Allowing the the fulfiller address to be separate from the rewardReceiver address.
 * @dev Entropy emitted by this contract is verified by the PlonkVerifier before being stored in the contract.
 * @dev To avoid brute forcing a desired outcome, the contract requires a predetermined commitment to be passed in by the fulfiller.
 * @dev This commitment is used as a public input to the PlonkVerifier and is used to verify the proof.
 * @dev The commitment is the hash of a secret that the fulfiller must use when generating the next entropy.
 * @dev Additionally, the seed of the request must always match the keccak256(requestId, msg.sender, blockhash(requestBlockSet[_requestId])) hash.
 * @dev The requestBlockSet is set to the blockNumber - 1 to ensure that the request is valid within the same block as the request.
 * @dev Since all public inputs are constrained, and the commitment is predetermined, we can trust that the entropy (if valid) is properly generated and not tampered with.
 */

contract FoundnoneVRF is PlonkVerifier, AccessControl {
    struct Request {
        address callbackAddress;
        uint32 callbackGasLimit;
        uint256 requestBlockSet;
        uint256 requestFeePaid;
        address requester;
    }

    // 5k is plenty for an EXTCODESIZE call (2600) + warm CALL (100)
    // and some arithmetic operations.
    uint256 private constant GAS_FOR_CALL_EXACT_CHECK = 5_000;

    /**
     * @notice The admin role for the contract
     * @dev The admin role is used to manage the request fee and contract fee percentage and to withdraw contract fees
     */
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    /**
     * @notice A mapping to track if a request has been fulfilled and to easily retrieve the entropy
     */
    mapping(uint256 => uint256) public entropies;

    /**
     * @notice The next requestId to be used
     * @dev This is incremented every time a new request is made
     */
    uint256 public nextRequestId = 0;

    /**
     * @notice The fee required to make a request
     */
    uint256 public requestFee = 0.000005 ether;

    /**
     * @notice The percentage of the fee that goes to the contract owner
     */
    uint256 public contractFeeBasisPoints = 500;

    /**
     * @notice A mapping to track the contract balance of each reward receiver
     */
    mapping(address => uint256) public rewardReceiverBalance;

    /**
     * @notice A mapping to track the parameters of each request
     */
    mapping(uint256 => Request) public request;

    /**
     * @notice A mapping to store the commitments for each fulfiller
     * @dev A commitment is a hash of a secret that the fulfiller must use when generating the next entropy
     * @dev When fulfilling a request, a new commitment should be passed to the contract to set the new commitment for that fulfiller
     * @dev This is used to ensure that fulfillers are not able to spoof entropies
     */
    mapping(address => uint256) public commitments;

    /**
     * @notice A mapping to track if a commitment is in use
     * @dev Ensures that two fulfillers cannot use the same commitment at the same time
     */
    mapping(uint256 => bool) public commitmentInUse;

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
     * @notice reentrancy lock active if true
     */
    bool reentrancyLock;

    /**
     * @notice Event emitted when a request is fulfilled, the fulfilling address is the fee receiver
     * @param requestId The requestId of the request
     * @param rewardReceiver The address to receive the reward
     * @param proof The proof data
     * @param publicInputs The public inputs of the proof
     */
    event RequestFulfilled(
        uint256 indexed requestId,
        bool callbackSuccess,
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
    event RngRequested(
        address callbackAddress,
        uint32 callbackGasLimit,
        uint256 requestId,
        bytes32 blockHash,
        uint256 requestBlockSet,
        address requester,
        uint256 feePaid
    );

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
    event RewardReceiverBalanceWithdrawn(
        address indexed rewardReceiver,
        uint256 amount
    );

    /**
     * @notice Event emitted when the request fee is updated
     * @param newFee The new fee
     */
    event RequestFeeUpdated(uint256 newFee);

    /**
     * @notice Event emitted when a request is refunded
     * @param requestId The requestId of the request
     * @param requester The address of the requester
     * @param amount The amount refunded
     */
    event RequestRefunded(
        uint256 indexed requestId,
        address indexed requester,
        uint256 amount
    );

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
    error InvalidSeedOrBlockHashUnavailable();
    error InvalidCommitmentBlock();
    error InsufficientFee();
    error RequestNotFulfilled();
    error InvalidFeeBasisPoints();
    error InsufficientBalance();
    error RequestStillValid();
    error InvalidRequester();
    error CommitmentInUse();
    error Reentrant();
    error RequesterNotWhitelisted();

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
     */
    function submitEntropy(
        uint256[24] calldata _proof,
        uint256[3] calldata _publicInputs,
        uint256 _requestId,
        address _rewardReceiver
    ) external nonReentrant {
        // verify the proof
        _verifyProof(_proof, _publicInputs, _requestId);

        // store the entropy
        entropies[_requestId] = _publicInputs[1];

        // calculate the fee
        uint256 fee = (requestFee * contractFeeBasisPoints) / 10_000;

        // update balances
        contractFeeBalance += fee;
        rewardReceiverBalance[_rewardReceiver] += requestFee - fee;

        bytes memory resp = abi.encodeWithSelector(
            IFoundnoneVRF.fulfillEntropyCallback.selector,
            _requestId,
            _publicInputs[1]
        );

        bool _success = false;

        if (
            request[_requestId].callbackAddress != address(0) &&
            request[_requestId].callbackGasLimit != 0
        ) {
            reentrancyLock = true;
            _success = _callWithExactGas(
                request[_requestId].callbackGasLimit,
                request[_requestId].callbackAddress,
                resp
            );
            reentrancyLock = false;
        }

        // emit the event
        emit RequestFulfilled(
            _requestId,
            _success,
            _rewardReceiver,
            _proof,
            _publicInputs
        );
    }

    /**
     * @dev calls target address with exactly gasAmount gas and data as calldata
     * or reverts if at least gasAmount gas is not available.
     */
    function _callWithExactGas(
        uint256 gasAmount,
        address target,
        bytes memory data
    ) private returns (bool success) {
        assembly {
            let g := gas()
            // Compute g -= GAS_FOR_CALL_EXACT_CHECK and check for underflow
            // The gas actually passed to the callee is min(gasAmount, 63//64*gas available).
            // We want to ensure that we revert if gasAmount >  63//64*gas available
            // as we do not want to provide them with less, however that check itself costs
            // gas.  GAS_FOR_CALL_EXACT_CHECK ensures we have at least enough gas to be able
            // to revert if gasAmount >  63//64*gas available.
            if lt(g, GAS_FOR_CALL_EXACT_CHECK) {
                revert(0, 0)
            }
            g := sub(g, GAS_FOR_CALL_EXACT_CHECK)
            // if g - g//64 <= gasAmount, revert
            // (we subtract g//64 because of EIP-150)
            if iszero(gt(sub(g, div(g, 64)), gasAmount)) {
                revert(0, 0)
            }
            // solidity calls check that a contract actually exists at the destination, so we do the same
            if iszero(extcodesize(target)) {
                revert(0, 0)
            }
            // call and return whether we succeeded. ignore return data
            // call(gas,addr,value,argsOffset,argsLength,retOffset,retLength)
            success := call(
                gasAmount,
                target,
                0,
                add(data, 0x20),
                mload(data),
                0,
                0
            )
        }
        return success;
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
        Request memory _request = request[_requestId];

        if (_requestId > nextRequestId) {
            revert InvalidRequestId();
        }
        if (entropies[_requestId] != 0) {
            revert RequestAlreadyFulfilled();
        }
        if (commitments[msg.sender] != (_publicInputs[2])) {
            revert InvalidCommitment();
        }
        if (commitmentBlockSet[msg.sender] > _request.requestBlockSet) {
            revert InvalidCommitmentBlock();
        }
        // since the blockhash is only available for the last 256 blocks, this fulfillment is inherently invalid if the request is older than 256 blocks
        uint256 _seedHash = uint256(
            keccak256(
                abi.encodePacked(
                    _requestId,
                    _request.requestBlockSet,
                    blockhash(_request.requestBlockSet)
                )
            )
        );
        // q is inherited from the PlonkVerifier contract and is the maximum value of the field
        uint256 _fieldHash = _seedHash % q;
        if (_publicInputs[0] != _fieldHash) {
            revert InvalidSeedOrBlockHashUnavailable();
        }
        if (!this.verifyProof(_proof, _publicInputs)) {
            revert InvalidProof();
        }
    }

    /**
     * @param _commitment The commitment to be set
     * @dev This function is used to set the commitment for the fulfiller
     * @dev The commitment is a hash of a secret that the fulfiller must use when generating the next entropy
     * @dev This is used to ensure that fulfillers are not able to spoof entropies
     * @dev The commitment is used as a public input to the PlonkVerifier and is used to verify the proof
     * @dev We also store the block number when the commitment was set to ensure that the commitment is only used on requests that are made after the commitment is set
     */
    function setCommitment(uint256 _commitment) public {
        if (commitmentInUse[_commitment]) {
            revert CommitmentInUse();
        }
        // unset the previous commitment
        if (commitments[msg.sender] != 0) {
            delete commitmentInUse[commitments[msg.sender]];
        }
        commitments[msg.sender] = _commitment;
        commitmentBlockSet[msg.sender] = block.number;
        commitmentInUse[_commitment] = true;
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
     * @dev We subtract 1 from the block to ensure a fulfillment is valid within the same block as the request.
     */
    function requestRng(
        address callbackAddress,
        uint32 callbackGasLimit
    ) external payable returns (uint256) {
        if (msg.value < requestFee) {
            revert InsufficientFee();
        }

        nextRequestId += 1;

        request[nextRequestId] = Request({
            callbackAddress: callbackAddress,
            callbackGasLimit: callbackGasLimit,
            requestBlockSet: block.number - 1, // we use the previous block to ensure the request is valid within the same block
            requestFeePaid: msg.value,
            requester: msg.sender
        });

        emit RngRequested(
            callbackAddress,
            callbackGasLimit,
            nextRequestId,
            blockhash(block.number - 1),
            block.number - 1,
            msg.sender,
            msg.value
        );

        return nextRequestId;
    }

    /**
     *
     * @param _requestId The requestId of the request
     * @dev This function is used to refund the request fee if the request is not fulfilled
     * @dev The requestId must be valid and the request must not have been fulfilled yet
     * @dev The blockhash of the request must be 0 to ensure that the request is not valid anymore
     * @dev The request fee is refunded to the requester (including the contract fee because it is not taken until fulfillment)
     */
    function refundUnfulfilledRequest(uint256 _requestId) external {
        Request storage _request = request[_requestId];
        if (_request.requester != msg.sender) {
            revert InvalidRequester();
        }
        if (entropies[_requestId] != 0) {
            revert RequestAlreadyFulfilled();
        }
        if (_request.requestBlockSet == 0) {
            revert InvalidRequestId();
        }
        if (blockhash(_request.requestBlockSet) != 0) {
            revert RequestStillValid();
        }
        payable(msg.sender).transfer(_request.requestFeePaid);
        delete request[_requestId];

        emit RequestRefunded(_requestId, msg.sender, _request.requestFeePaid);
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
    function withdrawContractFees() external onlyRole(ADMIN_ROLE) {
        uint256 balanceToWithdraw = contractFeeBalance;
        payable(msg.sender).transfer(balanceToWithdraw);
        contractFeeBalance = 0;
        emit ContractFeesWithdrawn(balanceToWithdraw);
    }

    /**
     * @param _newPercentage The new percentage to be set
     * @dev This function can only be called by the admin role
     */
    function setContractFeeBasisPoints(
        uint256 _newPercentage
    ) external onlyRole(ADMIN_ROLE) {
        if (_newPercentage > 2000) {
            revert InvalidFeeBasisPoints();
        }
        contractFeeBasisPoints = _newPercentage;
        emit ContractFeePercentageUpdated(_newPercentage);
    }

    modifier nonReentrant() {
        if (reentrancyLock) {
            revert Reentrant();
        }
        _;
    }
}
