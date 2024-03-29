// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProver.sol";
import "./interfaces/IFleetManagement.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract Prover is IProver, Ownable, ERC721 {
    event FleetManagementSetted(address indexed fleetManagement);

    uint256 nextId;

    mapping(uint256 => Type) _nodeType;
    mapping(uint256 => address) _operators;
    mapping(uint256 => bool) _paused;
    mapping(uint256 => PendingOperator) _pendingOperators;

    IFleetManagement public fleetManagement;

    modifier onlyProverOwner(uint256 _id) {
        require(ownerOf(_id) == msg.sender, "not owner");
        _;
    }

    constructor(string memory _name, string memory _symbol) ERC721(_name, _symbol) {}

    function nodeType(uint256 _id) external view override returns (Type) {
        _requireMinted(_id);
        return _nodeType[_id];
    }

    function operator(uint256 _id) external view override returns (address) {
        _requireMinted(_id);
        return _operators[_id];
    }

    function isPaused(uint256 _id) external view override returns (bool) {
        _requireMinted(_id);
        return _paused[_id];
    }

    function pendingOperator(uint256 _id) external view override returns (PendingOperator memory) {
        _requireMinted(_id);
        return _pendingOperators[_id];
    }

    function register(Type _type) external override returns (uint256 _id) {
        return register(_type, msg.sender);
    }

    function setFleetManagement(address _fleetManagement) external onlyOwner {
        require(_fleetManagement != address(0), "zero address");

        fleetManagement = IFleetManagement(_fleetManagement);
        emit FleetManagementSetted(_fleetManagement);
    }

    function register(Type _type, address _operator) public override returns (uint256 _id) {
        _id = ++nextId;

        _nodeType[_id] = _type;
        _operators[_id] = _operator;
        _paused[_id] = false;

        emit ProverCreated(_id, _operator);
    }

    function changeOperator(uint256 _id, address _operator) external override onlyProverOwner(_id) {
        require(_operator != address(0), "zero address");

        PendingOperator storage _pending = _pendingOperators[_id];
        _pending.timestamp = block.timestamp;
        _pending.operator = _operator;

        emit PendingOperatorAdded(_id, _operator);
    }

    function activePendingOperator(uint256 _id) external override {
        PendingOperator memory _pending = _pendingOperators[_id];

        require(
            _pending.timestamp > 0 && _pending.timestamp <= block.timestamp + fleetManagement.epoch(),
            "time to short"
        );
        _operators[_id] = _pending.operator;
        delete _operators[_id];

        emit OperatorActived(_id, _pending.operator);
    }

    function pause(uint256 _id) external override onlyProverOwner(_id) {
        require(!_paused[_id], "already paused");

        _paused[_id] = true;

        emit ProverPaused(_id);
    }

    function resume(uint256 _id) external override onlyProverOwner(_id) {
        require(_paused[_id], "already actived");

        _paused[_id] = false;

        emit ProverResumed(_id);
    }
}
