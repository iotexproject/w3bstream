// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IRouter} from "./interfaces/IRouter.sol";

contract W3bstreamRouter is IRouter {
    function registerProject(uint256 _projectId, address _verifier) external {}

    function submit(uint256 _projectId, bytes32 _prover, bytes calldata _data) external {}

    function receiver(uint256 _projectId) external view override returns (address) {}

    function fleetManager() external view override returns (address) {}

    function register(uint256 _projectId, address _receiver) external override {}

    function update(uint256 _projectId, address _receiver) external override {}
}
