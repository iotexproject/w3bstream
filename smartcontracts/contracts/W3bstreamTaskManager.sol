// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ITaskManager, TaskAssignment} from "./interfaces/ITaskManager.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

struct Record {
    bytes32 hash;
    address owner;
    address sequencer;
    address prover;
    uint256 rewardForProver;
    uint256 rewardForSequencer;
    uint256 deadline;
    bool settled;
}

interface IDebits {
    function withhold(address token, address owner, uint256 amount) external;
    function redeem(address token, address owner, uint256 amount) external;
    function distribute(address token, address owner, address[] calldata recipients, uint256[] calldata amounts) external;
}

interface IProjectReward {
    function isPaused(uint256 _id) external view returns (bool);
    function rewardToken(uint256 _id) external view returns (address);
    function rewardAmount(address owner, uint256 id) external view returns (uint256);
}

interface IProverStore {
    function isPaused(address prover) external view returns (bool);
    function rebateRatio(address prover) external view returns (uint16);
    function beneficiary(address prover) external view returns (address);
}

contract W3bstreamTaskManager is OwnableUpgradeable, ITaskManager {
    using ECDSA for bytes32;
    event TaskAssigned(uint256 indexed projectId, bytes32 indexed taskId, address indexed prover, uint256 deadline);
    event TaskSettled(uint256 indexed projectId, bytes32 indexed taskId, address prover);
    event OperatorAdded(address operator);
    event OperatorRemoved(address operator);
    mapping(uint256 => mapping(bytes32 => Record)) public records;
    mapping(address => bool) public operators;
    address public debits;
    address public projectReward;
    address public proverStore;

    modifier onlyOperator() {
        require(operators[msg.sender], "not operator");
        _;
    }

    function initialize(address _debits, address _projectReward) public initializer {
        __Ownable_init();
        debits = _debits;
        projectReward = _projectReward;
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
        bytes memory signature,
        address prover,
        uint256 deadline,
        address sequencer
    ) internal {
        IProverStore ps = IProverStore(proverStore);
        require(!ps.isPaused(prover), "prover paused");
        uint16 rebateRatio = ps.rebateRatio(prover);
        require(rebateRatio <= 10000, "invalid rebate ratio");
        IProjectReward pr = IProjectReward(projectReward);
        require(!pr.isPaused(projectId), "project paused");
        address taskOwner = hash.recover(signature);
        uint256 rewardAmount = pr.rewardAmount(taskOwner, projectId);
        address rewardToken = pr.rewardToken(projectId);
        IDebits(debits).withhold(rewardToken, taskOwner, rewardAmount);
        require(prover != address(0), "invalid prover");
        Record storage record = records[projectId][taskId];
        require(record.settled == false, "task already settled");
        if (record.prover != address(0)) {
            require(record.deadline < block.timestamp, "task already assigned");
        }
        record.hash = hash;
        record.owner = taskOwner;
        record.prover = prover;
        record.deadline = deadline;
        record.sequencer = sequencer;
        record.rewardForSequencer = rewardAmount * rebateRatio / 10000;
        record.rewardForProver = rewardAmount - record.rewardForSequencer;
        emit TaskAssigned(projectId, taskId, prover, deadline);
    }

    function assign(
        TaskAssignment calldata assignment,
        address sequencer,
        uint256 deadline
    ) public onlyOperator {
        _assign(assignment.projectId, assignment.taskId, assignment.hash, assignment.signature, assignment.prover, deadline, sequencer);
    }

    function assign(
        TaskAssignment[] calldata taskAssignments,
        address sequencer,
        uint256 deadline
    ) public onlyOperator {
        for (uint256 i = 0; i < taskAssignments.length; i++) {
            _assign(taskAssignments[i].projectId, taskAssignments[i].taskId, taskAssignments[i].hash, taskAssignments[i].signature, taskAssignments[i].prover, deadline, sequencer);
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
        address[] memory recipients = new address[](2);
        uint256[] memory amounts = new uint256[](2);
        recipients[0] = IProverStore(proverStore).beneficiary(record.prover);
        amounts[0] = record.rewardForProver;
        recipients[1] = record.sequencer;
        amounts[1] = record.rewardForSequencer;
        address rewardToken = IProjectReward(projectReward).rewardToken(projectId);
        IDebits(debits).distribute(rewardToken, record.owner, recipients, amounts);
    }

    function recall(uint256 projectId, bytes32 taskId) public {
        Record storage record = records[projectId][taskId];
        require(record.owner == msg.sender, "not owner");
        require(record.settled == false, "task already settled");
        require(record.deadline < block.number, "task assignement not expired");
        record.prover = address(0);
        record.deadline = 0;
        address rewardToken = IProjectReward(projectReward).rewardToken(projectId);
        IDebits(debits).redeem(rewardToken, record.owner, record.rewardForProver + record.rewardForSequencer);
    }

}
