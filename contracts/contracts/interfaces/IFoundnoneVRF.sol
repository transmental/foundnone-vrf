// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IFoundnoneVRF {
    function fulfillEntropyCallback(
        uint256 requestId,
        uint256[] calldata entropies
    ) external;

    function requestRng(
        address callbackAddress,
        uint32 callbackGasLimit
    ) external payable returns (uint256);
}
