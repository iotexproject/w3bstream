// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IDapp {
    function process(address _prover, uint256 _projectId, bytes32[] calldata _taskIds, bytes calldata _data) external;
}
