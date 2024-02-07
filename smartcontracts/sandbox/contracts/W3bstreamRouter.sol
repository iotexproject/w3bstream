// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {IFleetManager} from "./interfaces/IFleetManager.sol";
import {IRouter} from "./interfaces/IRouter.sol";

contract W3bstreamRouter is IRouter {
    address public override owner;
    address public override admin;
    address public override projectRegistry;
    address public override fleetManager;
    mapping(uint256 => address) public override receiver;

    constructor(address _projectRegistry, address _fleetManager) {
        owner = msg.sender;
        admin = msg.sender;
        projectRegistry = _projectRegistry;
        fleetManager = _fleetManager;
    }

    function register(uint256 _projectId, address _receiver) external override {
        if (_receiver == address(0)) revert ZeroAddress();
        if (IERC721(projectRegistry).ownerOf(_projectId) != msg.sender) {
            revert NotProjectOwner();
        }
        if (receiver[_projectId] != address(0)) revert AlreadyRegistered();

        receiver[_projectId] = _receiver;
        emit ProjectRegistered(_projectId, _receiver);
    }

    function submit(uint256 _projectId, bytes calldata _data) external override {
        address _receiver = receiver[_projectId];
        if (_receiver == address(0)) {
            revert UnregisterProject();
        }
        if (!IFleetManager(fleetManager).isAllowed(msg.sender, _projectId)) {
            revert NotOperator();
        }

        (bool success, ) = _receiver.call(_data);
        emit DataReceived(msg.sender, success);
    }

    function update(uint256 _projectId, address _receiver) external override {}

    function setFleetManager(address _fleetManager) external override {}

    function setOwner(address _owner) external override {}

    function setAdmin(address _admin) external override {}

    function setProjectRegistry(address _projectRegistry) external override {}
}
