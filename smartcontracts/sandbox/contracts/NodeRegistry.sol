// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {INodeRegistry} from "./interfaces/INodeRegistry.sol";

contract NodeRegistry is INodeRegistry, Initializable {
    // operator => node
    mapping(address => address) internal _nodes;
    // node => operator
    mapping(address => address) internal _operators;

    function initialize() public initializer {}

    function register(address _operator) public {
        if (_operators[msg.sender] != address(0)) {
            revert NodeAlreadyRegistered();
        }
        if (_nodes[_operator] != address(0)) {
            revert OperatorAlreadyRegistered();
        }

        _nodes[_operator] = msg.sender;
        _operators[msg.sender] = _operator;

        emit NodeRegistered(msg.sender, _operator);
    }

    function updateOperator(address _operator) public {
        if (_operators[msg.sender] == address(0)) {
            revert NodeUnregister();
        }
        if (_nodes[_operator] != address(0)) {
            revert OperatorAlreadyRegistered();
        }

        _operators[msg.sender] = _operator;
        emit NodeUpdated(msg.sender, _operator);
    }

    function getNode(address _operator) external view returns (Node memory) {
        address _node = _nodes[_operator];
        if (_node == address(0)) {
            revert NodeUnregister();
        }

        return Node(_node, _operator);
    }
}
