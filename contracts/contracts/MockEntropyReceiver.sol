// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./interfaces/IFoundnoneVRF.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

contract MockEntropyReceiver is AccessControl {
    IFoundnoneVRF public vrf;
    uint256 public latestRequestIdFulfilled;
    uint256 public latestEntropy;
    event CallbackReceived(uint256 indexed requestId, uint256 entropy);

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant VRF_ROLE = keccak256("VRF_ROLE");

    constructor(address vrfAddress) {
        vrf = IFoundnoneVRF(vrfAddress);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
        _grantRole(VRF_ROLE, vrfAddress);
    }

    function fulfillEntropyCallback(
        uint256 _requestId,
        uint256 _entropy
    ) external onlyRole(VRF_ROLE) {
        latestRequestIdFulfilled = _requestId;
        latestEntropy = _entropy;
        emit CallbackReceived(_requestId, latestEntropy);
    }

    function setVrfAddress(address vrfAddress) external onlyRole(ADMIN_ROLE) {
        require(vrfAddress != address(0), "Invalid VRF address");
        vrf = IFoundnoneVRF(vrfAddress);
    }
}
