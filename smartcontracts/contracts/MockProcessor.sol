// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MockProcessor {
    constructor() {}

    function process(uint256 _projectId, uint256 _proverId, string memory _clientId, bytes calldata _data) external {}
}
