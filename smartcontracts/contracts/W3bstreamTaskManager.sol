// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ITaskManager, TaskAssignment} from "./interfaces/ITaskManager.sol";

struct Record {
    bytes32 hash;
    address sequencer;
    address prover;
    uint256 rewardForProver;
    uint256 rewardForSequencer;
    uint256 deadline;
    bool settled;
}

/*
interface IDebits {
    function withhold(uint256 id, address owner, uint256 amount) external;
    function redeem(uint256 id, address owner, uint256 amount) external;
    function distribute(uint256 id, address owner, address[] calldata recipients, uint256[] calldata amounts) external;
}
*/
contract W3bstreamTaskManager is OwnableUpgradeable, ITaskManager {
    event TaskAssigned(uint256 indexed projectId, bytes32 indexed taskId, address indexed prover, uint256 deadline);
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
        bytes32 hash,
        address prover,
        uint256 deadline,
        address sequencer
    ) internal {
        require(prover != address(0), "invalid prover");
        Record storage record = records[projectId][taskId];
        require(record.settled == false, "task already settled");
        if (record.prover != address(0)) {
            require(record.deadline < block.timestamp, "task already assigned");
        }
        record.hash = hash;
        record.prover = prover;
        record.deadline = deadline;
        record.sequencer = sequencer;
        emit TaskAssigned(projectId, taskId, prover, deadline);
    }

    function assign(
        TaskAssignment calldata assignment,
        address sequencer,
        uint256 deadline
    ) public onlyOperator {
        _assign(assignment.projectId, assignment.taskId, assignment.hash, assignment.prover, deadline, sequencer);
    }

    function assign(
        TaskAssignment[] calldata taskAssignments,
        address sequencer,
        uint256 deadline
    ) public onlyOperator {
        for (uint256 i = 0; i < taskAssignments.length; i++) {
            _assign(taskAssignments[i].projectId, taskAssignments[i].taskId, taskAssignments[i].hash, taskAssignments[i].prover, deadline, sequencer);
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
