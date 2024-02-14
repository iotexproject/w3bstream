// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @title IWSRouter
/// @notice W3bstream router interface.
interface IWSRouter {
    event DataReceived(address indexed operator, bool success, string revertReason);
    event FleetManagerChanged(address indexed fleetManager);
    event OwnerChanged(address indexed owner);
    event AdminChanged(address indexed admin);
    event ProjectRegistryChanged(address indexed projectRegistry);

    error NotProjectOwner();
    error NotOperator();
    error NotOwner();
    error NotAdmin();
    error ZeroAddress();

    /// @notice project regsitry contract
    /// @return address of project regsitry
    function projectRegistry() external view returns (address);

    /// @notice fleet manager
    /// @return address of fleet manager
    function fleetManager() external view returns (address);

    /// @notice submit data to project
    /// @param _projectId project id
    /// @param _receiver project data reveiver
    /// @param _tunnelId tunnel id
    /// @param _batchMR batch merkle root
    /// @param _devicesMR devices merkle root
    /// @param _zkProof zk proof
    function submit(
        uint256 _projectId,
        address _receiver,
        uint256 _tunnelId,
        bytes32 _batchMR,
        bytes32 _devicesMR,
        bytes calldata _zkProof
    ) external;

    /// @notice set project registry
    /// @param _projectRegistry project registry
    function setProjectRegistry(address _projectRegistry) external;

    /// @notice set fleet manager
    /// @param _fleetManager fleet manager
    function setFleetManager(address _fleetManager) external;

    /// @notice router owner
    /// @return router owner
    function owner() external view returns (address);

    /// @notice set router owner
    /// @param _owner .
    function setOwner(address _owner) external;

    /// @notice router admin
    /// @return router admin`
    function admin() external view returns (address);

    /// @notice set router admin
    /// @param _admin .
    function setAdmin(address _admin) external;
}
