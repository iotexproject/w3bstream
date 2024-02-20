// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @title Sandbox fleet manager
interface IFleetManager {
    event NodeAllowed(uint256 indexed projectId, address indexed node);
    event NodeDisallowed(uint256 indexed projectId, address indexed node);

    function allow(uint256 _projectId, address _node) external;

    function disallow(uint256 _projectId, address _node) external;

    /// @notice check operator rights for project
    /// @param _operator operator address
    /// @param _projectId project id
    function isAllowed(address _operator, uint256 _projectId) external view returns (bool);

    error NotProjectOwner();
    error NodeAlreadyAllowed();
    error NodeNotAllow();
    error NodeUnregister();
    error InvalidOperatorAddress();
}
