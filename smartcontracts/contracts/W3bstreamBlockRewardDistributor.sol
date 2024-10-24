// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract W3bstreamBlockRewardDistributor is OwnableUpgradeable {
    event Distributed(address indexed recipient, uint256 amount);
    event Withdrawn(uint256 amount);
    event OperatorSet(address indexed operator);
    event Topup(uint256 amount);
    address public operator;

    modifier onlyOperator() {
        require(msg.sender == operator, "not block reward distributor operator");
        _;
    }

    receive() external payable {
        emit Topup(msg.value);
    }

    function initialize() public initializer {
        __Ownable_init();
    }

    function setOperator(address _operator) public onlyOwner {
        operator = _operator;
        emit OperatorSet(_operator);
    }

    function distribute(address recipient, uint256 amount) public onlyOperator {
        if (amount == 0) {
            return;
        }
        if (amount > address(this).balance) {
            revert("insufficient balance");
        }
        (bool success, ) = recipient.call{value: amount}("");
        require(success, "transfer failed");
        emit Distributed(recipient, amount);
    }

    function withdraw(uint256 amount) public onlyOwner {
        require(amount <= address(this).balance, "insufficient balance");
        (bool success, ) = msg.sender.call{value: amount}("");
        require(success, "transfer failed");
        emit Withdrawn(amount);
    }
}
