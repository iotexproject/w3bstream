// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IDapp {
    function process(
        uint256 _projectId,
        bytes32 _taskId,
        address _prover,
        address _deviceId,
        bytes calldata _data
    ) external;
}
