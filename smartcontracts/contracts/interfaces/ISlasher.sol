// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface ISlasher {
    function slash(address prover, uint256 amount) external;
}
