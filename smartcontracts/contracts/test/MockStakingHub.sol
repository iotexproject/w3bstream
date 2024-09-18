// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MockStakingHub {
    uint256 public amount;

    constructor(uint256 _amount) {
        amount = _amount;
    }

    function setAmount(uint256 _amount) external {
        amount = _amount;
    }

    function stakedAmount(address) external view returns (uint256) {
        return amount;
    }
}
