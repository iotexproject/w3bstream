// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract OperatorRegistry {
    uint256 private _nextOperatorId;

    mapping(uint256 => Operator) public operators;

    struct Operator {
        // an uri can be added with operator metadata
        // name can be in metadata
        string name;
        address profile;
        address node;
        address rewards;
    }

    event OperatorRegistered(
        uint256 indexed operatorId,
        address indexed profile,
        address indexed node,
        address rewards,
        string name
    );
    event OperatorNodeUpdated(uint256 indexed operatorId, address indexed newNode);
    event OperatorRewardsUpdated(uint256 indexed operatorId, address indexed rewards);

    modifier onlyExistingOperator(uint256 operatorId) {
        require(operators[operatorId].profile != address(0), "OperatorRegistry: unexistent operator");
        _;
    }

    modifier onlyOperatorOwner(uint256 operatorId) {
        require(operators[operatorId].profile == msg.sender, "OperatorRegistry: Not operator owner");
        _;
    }

    function registerOperator(Operator memory operator) public {
        require(
            keccak256(abi.encodePacked(operator.name)) != keccak256(abi.encodePacked("")),
            "OperatorRegistry: name can't be empty"
        );

        uint256 operatorId = _nextOperatorId++;

        // Should we extract profileAddress from msg.sender?
        operators[operatorId] = operator;

        emit OperatorRegistered(operatorId, operator.profile, operator.node, operator.rewards, operator.name);
    }

    function updateNode(
        uint256 operatorId,
        address newNode
    ) public onlyExistingOperator(operatorId) onlyOperatorOwner(operatorId) {
        operators[operatorId].node = newNode;

        emit OperatorNodeUpdated(operatorId, newNode);
    }

    function updateRewards(
        uint256 operatorId,
        address newRewards
    ) public onlyExistingOperator(operatorId) onlyOperatorOwner(operatorId) {
        operators[operatorId].rewards = newRewards;

        emit OperatorRewardsUpdated(operatorId, newRewards);
    }

    // function stake() public {}
    // function unstake() public {}
}
