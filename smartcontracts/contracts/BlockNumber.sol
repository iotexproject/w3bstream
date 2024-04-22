// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;


contract BlockNumber {
    function blockNumber() external view returns (uint256) {
       return block.number;
    }
}
