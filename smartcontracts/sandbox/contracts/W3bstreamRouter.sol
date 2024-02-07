// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IRouter} from "./interfaces/IRouter.sol";

contract W3bstreamRouter is IRouter {
    address public override owner;
    address public override admin;

    function register(uint256 _projectId, address _receiver) external override {}

    function submit(uint256 _projectId, bytes32 _prover, bytes calldata _data) external override {}

    function receiver(uint256 _projectId) external view override returns (address) {}

    function fleetManager() external view override returns (address) {}

    function update(uint256 _projectId, address _receiver) external override {}

    function setFleetManager(address _fleetManager) external override {}

    function setOwner(address _owner) external override {}

    function setAdmin(address _admin) external override {}
}
