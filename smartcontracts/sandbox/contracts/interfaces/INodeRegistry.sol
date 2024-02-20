// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @notice Web3stream node registry
interface INodeRegistry {
    struct Node {
        address node;
        address operator;
        // add stake field future
    }

    error NodeAlreadyRegistered();
    error OperatorAlreadyRegistered();
    error NodeUnregister();

    event NodeRegistered(address indexed node, address indexed operator);
    event NodeUpdated(address indexed node, address indexed operator);

    /// @notice get Node by operator address
    /// @param _operator operator address
    function getNode(address _operator) external view returns (Node memory);
}
