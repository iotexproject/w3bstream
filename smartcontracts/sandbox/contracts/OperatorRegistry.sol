// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract OperatorRegistry {
    mapping(address => Operator) public operators;

    struct Operator {
        address node;
        address rewards;
    }

    event OperatorRegistered(address indexed profile, address indexed node, address rewards);
    event OperatorNodeUpdated(address indexed profile, address indexed newNode);
    event OperatorRewardsUpdated(address indexed profile, address indexed rewards);

    modifier onlyExistingOperator() {
        require(operators[msg.sender].node != address(0), "OperatorRegistry: unexistent operator");
        _;
    }

    modifier onlyNewOperator() {
        require(operators[msg.sender].node == address(0), "OperatorRegistry: operator already registered");
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

    // function stake() public {}
    // function unstake() public {}
}
