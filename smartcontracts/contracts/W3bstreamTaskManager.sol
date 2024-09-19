// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

struct Record {
    address prover;
    uint256 deadline;
    bool settled;
}

struct TaskAssignment {
    uint256 projectId;
    bytes32 taskId;
    address prover;
}

contract W3bstreamTaskManager is OwnableUpgradeable {
    event TaskAssigned(uint256 indexed projectId, bytes32 indexed taskId, address prover, uint256 deadline);
    event TaskSettled(uint256 indexed projectId, bytes32 indexed taskId, address prover);
    event OperatorAdded(address operator);
    event OperatorRemoved(address operator);
    mapping(uint256 => mapping(bytes32 => Record)) public records;
    mapping(address => bool) public operators;

    modifier onlyOperator() {
        require(operators[msg.sender], "not operator");
        _;
    }

    function initialize() public initializer {
        __Ownable_init();
    }

    function addOperator(address operator) public onlyOwner {
        require(operator != address(0), "invalid operator");
        require(!operators[operator], "operator already added");
        operators[operator] = true;
        emit OperatorAdded(operator);
    }

    function removeOperator(address operator) public onlyOwner {
        require(operators[operator], "operator not found");
        delete operators[operator];
        emit OperatorRemoved(operator);
    }

    function _assign(
        uint256 projectId,
        bytes32 taskId,
        address prover,
        uint256 deadline
    ) internal {
        require(prover != address(0), "invalid prover");
        Record storage record = records[projectId][taskId];
        require(record.settled == false, "task already settled");
        if (record.prover != address(0)) {
            require(record.deadline < block.timestamp, "task already assigned");
        }
        record.prover = prover;
        record.deadline = deadline;
        emit TaskAssigned(projectId, taskId, prover, deadline);
    }

    function assign(
        TaskAssignment calldata assignment,
        uint256 deadline
    ) public onlyOperator {
        _assign(assignment.projectId, assignment.taskId, assignment.prover, deadline);
    }

    function assign(
        TaskAssignment[] calldata taskAssignments,
        uint256 deadline
    ) public onlyOperator {
        for (uint256 i = 0; i < taskAssignments.length; i++) {
            _assign(taskAssignments[i].projectId, taskAssignments[i].taskId, taskAssignments[i].prover, deadline);
        }
    }

    function settle(uint256 projectId, bytes32 taskId, address prover) public onlyOperator {
        require(prover != address(0), "invalid prover");
        Record storage record = records[projectId][taskId];
        require(record.prover == prover, "invalid prover");
        require(record.deadline >= block.timestamp, "task assignement expired");
        require(record.settled == false, "task already settled");
        record.settled = true;
        emit TaskSettled(projectId, taskId, prover);
        // TODO: distribute task reward
    }

}
