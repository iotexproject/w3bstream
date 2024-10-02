// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface ITaskManager {
    function assign(TaskAssignment[] calldata assignments, address sequencer, uint256 deadline) external;
}

struct TaskAssignment {
    uint256 projectId;
    bytes32 taskId;
    bytes32 hash;
    bytes signature;
    address prover;
}
