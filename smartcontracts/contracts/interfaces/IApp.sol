// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IApp {
    function process(bytes calldata _data) external;
}
