// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IOperatorRegistry} from "./interfaces/IOperatorRegistry.sol";

contract OperatorRegistry is IOperatorRegistry {
    mapping(address => Operator) internal operators;

    event OperatorRegistered(address indexed profile, address indexed node, address rewards);
    event OperatorNodeUpdated(address indexed profile, address indexed newNode);
    event OperatorRewardsUpdated(address indexed profile, address indexed rewards);

    modifier onlyExistingOperator() {
        if (operators[msg.sender].node == address(0)) {
            revert UnexistentOperator();
        }
        _;
    }

    modifier onlyNewOperator() {
        if (operators[msg.sender].node != address(0)) {
            revert OperatorAlreadyRegistered();
        }
        _;
    }

    function registerOperator(Operator memory _operator) public onlyNewOperator {
        address profile = msg.sender;
        operators[profile] = _operator;
        emit OperatorRegistered(profile, _operator.node, _operator.rewards);
    }

    function updateNode(address _newNode) public onlyExistingOperator {
        address profile = msg.sender;
        operators[profile].node = _newNode;
        emit OperatorNodeUpdated(profile, _newNode);
    }

    function updateRewards(address _newRewards) public onlyExistingOperator {
        address profile = msg.sender;
        operators[profile].rewards = _newRewards;
        emit OperatorRewardsUpdated(profile, _newRewards);
    }

    function getOperator(address _profile) external view returns (Operator memory) {
        return operators[_profile];
    }

    // function stake() public {}
    // function unstake() public {}
}
