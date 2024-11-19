// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract MockIoIDRegistry is OwnableUpgradeable {
    mapping(address => uint256) ioids;

    address public ioid;

    function initialize(address _ioid) public initializer {
        __Ownable_init();
        ioid = _ioid;
    }

    function bind(uint256 _ioid, address _device) external {
        ioids[_device] = _ioid;
    }

    function deviceTokenId(address _device) external view returns (uint256) {
        return ioids[_device];
    }

    function ioID() external view returns (address) {
        return ioid;
    }
}
