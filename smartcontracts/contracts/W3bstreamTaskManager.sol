// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

struct TaskAssignment {
    address prover;
    uint256 deadline;
    bool settled;
}

contract W3bstreamTaskManager is OwnableUpgradeable {
    event TaskAssigned(uint64 indexed projectId, uint64 indexed taskId, address prover, uint256 deadline);
    event TaskSettled(uint64 indexed projectId, uint64 indexed taskId, address prover);
    event OperatorAdded(address operator);
    event OperatorRemoved(address operator);
    mapping(uint64 => mapping(uint64 => TaskAssignment)) public assignments;
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

    function assign(
        uint64 projectId,
        uint64 taskId,
        address prover,
        uint256 deadline
    ) public onlyOperator{
        require(prover != address(0), "invalid prover");
        TaskAssignment storage assignment = assignments[projectId][taskId];
        require(assignment.settled == false, "task already settled");
        if (assignment.prover != address(0)) {
            require(assignment.deadline < block.timestamp, "task already assigned");
        }
        assignment.prover = prover;
        assignment.deadline = deadline;
        emit TaskAssigned(projectId, taskId, prover, deadline);
    }

    function settle(uint64 projectId, uint64 taskId, address prover) public onlyOperator {
        require(prover != address(0), "invalid prover");
        TaskAssignment storage assignment = assignments[projectId][taskId];
        require(assignment.prover == prover, "invalid prover");
        require(assignment.deadline >= block.timestamp, "task assignement expired");
        require(assignment.settled == false, "task already settled");
        assignment.settled = true;
        emit TaskSettled(projectId, taskId, prover);
        // TODO: distribute task reward
    }

}
