// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./interfaces/IFoundnoneVRF.sol";

contract MockEntropyReceiver {
    IFoundnoneVRF public vrf;
    uint256 public latestRequestIdFulfilled;
    uint256 public latestEntropy;
    event CallbackReceived(uint256 indexed requestId, uint256 entropy);

    constructor(address vrfAddress) {
        vrf = IFoundnoneVRF(vrfAddress);
    }

    function fulfillEntropyCallback(
        uint256 _requestId,
        uint256 _entropy
    ) external {
        require(msg.sender == address(vrf), "only vrf");
        latestRequestIdFulfilled = _requestId;
        latestEntropy = _entropy;
        emit CallbackReceived(_requestId, latestEntropy);
    }
}
