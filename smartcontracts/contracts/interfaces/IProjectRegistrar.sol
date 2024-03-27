// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IProjectRegistrar {
    struct Project {
        bool paused;
        address operator;
        bytes32 hash;
        string uri;
        mapping(bytes32 => bytes) metadata;
    }

    event FeeChanged(uint256 indexed fee);

    event ProjectCreated(address indexed owner, uint256 indexed projectId);
    event ProjectOperatorChanged(uint256 indexed projectId, address indexed operator);
    event ProjectPaused(uint256 indexed projectId);
    event ProjectResumed(uint256 indexed projectId);
    event ProjectConfigUpdated(uint256 indexed projectId);
    event ProjectMetadataUpdated(uint256 indexed projectId, bytes32 name);
    event ProjectMetadataRemoved(uint256 indexed projectId, bytes32 name);

    function fee() external view returns (uint256);

    function isPaused(uint256 _projectId) external view returns (bool);
    function operator(uint256 _projectId) external view returns (address);
    function hash(uint256 _projectId) external view returns (bytes32);
    function uri(uint256 _projectId) external view returns (string calldata);
    function metadata(uint256 _projectId, bytes32 _name) external view returns (bytes calldata);

    function register(string calldata _uri, bytes32 _hash) external payable returns (uint256 _proejctId);
    function changeOperator(uint256 _projectId, address _operator) external;
    function pause(uint256 _projectId) external;
    function resume(uint256 _projectId) external;
    function updateConfig(uint256 _projectId, string calldata _uri, bytes32 _hash) external;
    function updateMetadata(uint256 _projectId, bytes32 _name, bytes calldata _value) external;
    function removeMetadata(uint256 _projectId, bytes32 _name) external;
}
