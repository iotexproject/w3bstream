// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IRouter {
    function route(uint256 _projectId, address _app, bytes calldata _data) external;
}
