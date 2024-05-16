// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IioID {
    function ownerOf(uint256 tokenId) external view returns (address owner);
}

interface IioIDRegistry {
    function ioID() external view returns (address);
    function deviceTokenId(address device) external view returns (uint256);
}

interface IW3bstreamProject {
    function isValidProject(uint256 _projectId) external view returns (bool);
}

contract ProjectDevice {
    event Approve(uint256 projectId, address indexed device);
    event Unapprove(uint256 projectId, address indexed device);

    IioIDRegistry public ioIDRegistry;
    IW3bstreamProject public w3bstreamProject;

    mapping(uint256 => mapping(address => bool)) devices;

    constructor(address _ioIDRegistry, address _w3bstreamProject) {
        ioIDRegistry = IioIDRegistry(_ioIDRegistry);
        w3bstreamProject = IW3bstreamProject(_w3bstreamProject);
    }

    function approve(uint256 _projectId, address _device) external {
        require(w3bstreamProject.isValidProject(_projectId), "invalid project");
        require(!devices[_projectId][_device], "already approved");

        uint256 _tokenId = ioIDRegistry.deviceTokenId(_device);
        IioID ioID = IioID(ioIDRegistry.ioID());
        require(ioID.ownerOf(_tokenId) == msg.sender, "not device owner");

        devices[_projectId][_device] = true;
        emit Approve(_projectId, _device);
    }

    function unapprove(uint256 _projectId, address _device) external {
        require(devices[_projectId][_device], "not approve");
        require(w3bstreamProject.isValidProject(_projectId), "invalid project");

        uint256 _tokenId = ioIDRegistry.deviceTokenId(_device);
        IioID ioID = IioID(ioIDRegistry.ioID());
        require(ioID.ownerOf(_tokenId) == msg.sender, "not device owner");

        devices[_projectId][_device] = false;
        emit Unapprove(_projectId, _device);
    }

    function approved(uint256 _projectId, address _device) external view returns (bool) {
        return devices[_projectId][_device];
    }
}
