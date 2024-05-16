// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IioID {
    function ownerOf(uint256 tokenId) external view returns (address owner);
}

interface IioIDRegistry {
    function ioID() external view returns (address);
    function deviceTokenId(address device) external view returns (uint256);
}

contract ProjectDevice {
    event Approve(uint256 projectId, address indexed device);
    event Unapprove(uint256 projectId, address indexed device);

    IioIDRegistry public ioIDRegistry;
    uint256 public projectId;

    mapping(address => bool) devices;

    constructor(address _ioIDRegistry, uint256 _projectId) {
        ioIDRegistry = IioIDRegistry(_ioIDRegistry);
        projectId = _projectId;
    }

    function approve(address _device) external {
        require(!devices[_device], "already approved");

        uint256 _tokenId = ioIDRegistry.deviceTokenId(_device);
        IioID ioID = IioID(ioIDRegistry.ioID());
        require(ioID.ownerOf(_tokenId) == msg.sender, "not device owner");

        devices[_device] = true;
        emit Approve(projectId, _device);
    }

    function unapprove(address _device) external {
        require(devices[_device], "not approve");

        uint256 _tokenId = ioIDRegistry.deviceTokenId(_device);
        IioID ioID = IioID(ioIDRegistry.ioID());
        require(ioID.ownerOf(_tokenId) == msg.sender, "not device owner");

        devices[_device] = false;
        emit Unapprove(projectId, _device);
    }

    function approved(address _device) external view returns (bool) {
        return devices[_device];
    }
}
