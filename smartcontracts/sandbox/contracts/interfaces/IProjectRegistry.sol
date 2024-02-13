// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IProjectRegistry {
    struct Project {
        string uri;
        bytes32 hash;
        bool paused;
    }

    function getProject(uint256 _projectId) external view returns (Project memory);

    function isProjectOwner(address _account, uint256 _projectId) external view returns (bool);

    error OnlyOwnerAllowed();
    error EmptyUriValue();
    error ProjectAlreadyPaused();
    error ProjectNotPaused();
}
