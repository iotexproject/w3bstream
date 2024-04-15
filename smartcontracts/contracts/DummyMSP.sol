// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract DummyMSP {
    uint256 amount;
    constructor() {}

    function setAmount(uint256 _amount) external {
        amount = _amount;
    }

    function stakedAmount(address) external view returns (uint256) {
        return amount;
    }
}