// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IProjectStore {
    struct Project {
        string uri;
        bytes32 hash;
    }

    event AttributeSet(uint256 indexed projectId, bytes32 indexed key, bytes value);
    event AttributeRemoved(uint256 indexed projectId, bytes32 indexed key);
    event ProjectPaused(uint256 indexed projectId);
    event ProjectResumed(uint256 indexed projectId);
    event ProjectUpdated(uint256 indexed projectId, string uri, bytes32 hash);

    function isPaused(uint256 _projectId) external view returns (bool);
    function project(uint256 _projectId) external view returns (Project calldata);
    function attribute(uint256 _projectId, bytes32 _name) external view returns (bytes calldata);
    function attributesOf(uint256 _projectId, bytes32[] calldata _keys) external view returns (bytes[] calldata);
    function count() external view returns (uint256);

    function mint(address owner, string calldata _uri, bytes32 _hash) external returns (uint256 _projectId);
    function pause(uint256 _projectId) external;
    function resume(uint256 _projectId) external;
    function updateProject(uint256 _projectId, string memory _uri, bytes32 _hash) external;
}
