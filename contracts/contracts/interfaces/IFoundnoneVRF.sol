// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IFoundnoneVRF {
    function fulfillEntropyCallback(
        uint256 requestId,
        uint256[] calldata entropies
    ) external;
}
