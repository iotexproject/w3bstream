// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract W3bstreamBlockRewardDistributor is OwnableUpgradeable {
    address public operator;

    modifier onlyOperator() {
        require(msg.sender == operator, "not operator");
        _;
    }

    function initialize() public initializer {
        __Ownable_init();
    }

    function setOperator(address _operator) public onlyOwner {
        operator = _operator;
    }

    function distribute(address recipient, uint256 amount) public onlyOperator {
        (bool success, ) = recipient.call{value: amount}("");
        require(success, "transfer failed");
    }
}
