// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IDeviceRegistry {
    struct Device {
        string uri;
        bool enabled;
        bytes32 pubKey;
    }

    // @dev Will be run by the sprout node to get the device pubKey and status
    // @param _deviceId The device NFT id
    function getDeviceInfo(uint256 _deviceId) external view returns (Device memory);
}
