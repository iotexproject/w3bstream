// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @title Sandbox fleet manager
interface IFleetManager {
    /// @notice check operator rights for project
    /// @param _operator operator address
    /// @param _projectId project id
    function isAllowed(address _operator, uint256 _projectId) external view returns (bool);
}
