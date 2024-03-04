// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @notice Web3stream node registry
interface INodeRegistry {
    struct Node {
        bool active;
        address operator;
        // add stake field future
    }

    event NodeRegistered(address indexed owner, uint256 indexed nodeId, address indexed operator);
    event NodeUpdated(uint256 indexed nodeId, address indexed operator);

    error InvalidAddress();
    error NotNodeOwner();
    error OperatorAlreadyRegistered();
    error OperatorUnregister();

    function register(address _operator) external;

    function updateOperator(uint256 _tokenId, address _operator) external;

    function getNode(uint256 _tokenId) external view returns (Node memory);

    function getNodeAddress(uint256 _tokenId) external view returns (address);

    function getNodeByOperator(address _operator) external view returns (uint256 _nodeId, Node memory _node);
}
