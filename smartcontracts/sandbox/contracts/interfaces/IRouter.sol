// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @title IRouter
/// @notice W3bstream router interface.
interface IRouter {
    event ReceiverRegistered(uint256 indexed projectId, address indexed receiver);
    event ReceiverUnregistered(uint256 indexed projectId, address indexed receiver);
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
    error AlreadyRegistered();
    error ReceiverUnregister();

    /// @notice check project receiver
    /// @param _projectId project id
    /// @return project receiver address
    function isReceiver(uint256 _projectId, address _receiver) external view returns (bool);

    /// @notice project regsitry contract
    /// @return address of project regsitry
    function projectRegistry() external view returns (address);

    /// @notice fleet manager
    /// @return address of fleet manager
    function fleetManager() external view returns (address);

    /// @notice register project data receiver
    /// @param _projectId project id
    /// @param _receiver project data reveiver
    function register(uint256 _projectId, address _receiver) external;

    /// @notice unregister project data receiver
    /// @param _projectId project id
    /// @param _receiver project data reveiver
    function unregister(uint256 _projectId, address _receiver) external;

    /// @notice submit data to project
    /// @param _projectId project id
    /// @param _receiver project data reveiver
    /// @param _data data
    function submit(uint256 _projectId, address _receiver, bytes calldata _data) external;

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
