// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IProver {
    enum Type {
        General,
        Dephy
    }

    struct PendingOperator {
        uint256 timestamp;
        address operator;
    }

    event ProverCreated(uint256 indexed id, address indexed operator);
    event PendingOperatorAdded(uint256 indexed id, address indexed operator);
    event OperatorActived(uint256 indexed id, address indexed operator);
    event ProverPaused(uint256 indexed _id);
    event ProverResumed(uint256 indexed _id);

    function nodeType(uint256 _id) external view returns (Type);
    function operator(uint256 _id) external view returns (address);
    function isPaused(uint256 _id) external view returns (bool);
    function pendingOperator(uint256 _id) external view returns (PendingOperator calldata);

    function register(Type _type) external returns (uint256);
    function register(Type _type, address _operator) external returns (uint256);
    function changeOperator(uint256 _id, address _operator) external;
    function activePendingOperator(uint256 _id) external;
    function pause(uint256 _id) external;
    function resume(uint256 _id) external;
}
