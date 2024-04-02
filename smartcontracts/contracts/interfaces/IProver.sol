// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IProver {
    // TODO: rename file, add timestamp to pause/unpause
    struct PendingOperator {
        uint256 timestamp;
        address operator;
    }

    event PendingOperatorAdded(uint256 indexed id, address indexed operator);
    event NodeTypeUpdated(uint256 indexed id, uint256 type);
    event ProverPaused(uint256 indexed _id);
    event ProverResumed(uint256 indexed _id);

    function count() external view returns (uint256);
    function nodeType(uint256 _id) external view returns (uint256);
    function operator(uint256 _id) external view returns (address);
    function isPaused(uint256 _id) external view returns (bool);
    function pendingOperator(uint256 _id) external view returns (PendingOperator calldata);

    function register(uint256 _type) external returns (uint256);
    function register(uint256 _type, address _operator) external returns (uint256);
    function changeOperator(uint256 _id, address _operator) external;
    function updateNodeType(uint256 _id, uint256 _type) external;
    function pause(uint256 _id) external;
    function resume(uint256 _id) external;
}
