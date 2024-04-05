// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProver.sol";

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

contract Prover is IProver, OwnableUpgradeable, ERC721Upgradeable {
    uint256 nextId;

    mapping(uint256 => uint256) _nodeTypes;
    mapping(uint256 => address) _operators;
    mapping(uint256 => bool) _paused;
    mapping(address => uint256) operatorToProver;
    address minter;

    modifier onlyProverOwner(uint256 _id) {
        require(ownerOf(_id) == msg.sender, "not owner");
        _;
    }

    function initialize(string memory _name, string memory _symbol) public initializer {
        __Ownable_init();
        __ERC721_init(_name, _symbol);
    }

    function count() external view override returns (uint256) {
        return nextId + 1;
    }

    function nodeType(uint256 _id) external view override returns (uint256) {
        _requireMinted(_id);
        return _nodeTypes[_id];
    }

    function operator(uint256 _id) external view override returns (address) {
        _requireMinted(_id);
        return _operators[_id];
    }

    function ownerOfOperator(address _operator) external view override returns (uint256, address) {
        uint256 id = operatorToProver[_operator];
        require(id != 0, "invalid operator");
        return (id, ownerOf(id));
    }

    function isPaused(uint256 _id) external view override returns (bool) {
        _requireMinted(_id);
        return _paused[_id];
    }

    function register(Type _type, address _operator) public override returns (uint256 _id) {
        require(msg.sender == minter, "not minter");
        _id = ++nextId;
        _mint(_operator, _id);

        _paused[_id] = true;
        updateNodeTypeInternal(_id, _type);
        updateOperatorInternal(_id, _operator);
    }

    function updateOperatorInternal(uint256 _id, address _operator) internal {
        require(operatorToProver[_operator] != _id, "duplicate operator");
        operatorToProver[_operator] = _id;
        emit OperatorSet(_id, _operator);
    }

    function updateNodeTypeInternal(uint256 _id, uint256 _type) internal {
        _nodeTypes[_id] = _type;
        emit NodeTypeUpdated(_id, _type);
    }

    function updateNodeType(uint256 _id, uint256 _type) external override onlyProverOwner(_id) {
        updateNodeTypeInternal(_id, _type);
    }

    function changeOperator(uint256 _id, address _operator) external override onlyProverOwner(_id) {
        require(_operator != address(0), "zero address");
        PendingOperator memory pending = _pendingOperators[_id];
        if (pending.timestamp > 0 && pending.timestamp >= block.timestamp) {
            _operators[_id] = pending.operator;
        }
        updateOperatorInternal(_id, _operator);
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

    function changeMinter(address _minter) external onlyOwner {
        minter = _minter;

        emit MinterChanged(_minter);
    }
}
