// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IOperatorRegistry {
    struct Operator {
        address node;
        address rewards;
    }

    function getOperator(address) external view returns (Operator memory);
}
