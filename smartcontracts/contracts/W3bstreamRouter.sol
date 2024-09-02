// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IRouter.sol";
import "./interfaces/IFleetManagement.sol";
import "./interfaces/IDapp.sol";

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

interface IProjectCenter {
    function isPaused(uint256 _projectId) external view returns (bool);
}

contract W3bstreamRouter is IRouter, Initializable {
    address public fleetManagement;
    address public projectStore;

    mapping(uint256 => address) public override dapp;

    modifier onlyProjectOwner(uint256 _projectId) {
        require(IERC721(projectStore).ownerOf(_projectId) == msg.sender, "not project owner");
        _;
    }

    function initialize(address _fleetManagement, address _projectStore) public initializer {
        fleetManagement = _fleetManagement;
        projectStore = _projectStore;
    }

    function route(uint256 _projectId, uint256 _proverId, string memory _clientId, bytes calldata _data) external override {
        address _dapp = dapp[_projectId];
        require(_dapp != address(0), "no dapp");
        IFleetManagement _fm = IFleetManagement(fleetManagement);
        require(_fm.isActiveCoordinator(msg.sender, _projectId), "invalid coordinator");
        // TODO: 1. epoch based
        // TODO: 2. validate operator (of prover) signature
        require(_fm.isActiveProver(_proverId), "invalid prover");
        require(!IProjectCenter(projectStore).isPaused(_projectId), "invalid project");

        try IDapp(_dapp).process(_projectId, _proverId, _clientId, _data) {
            _fm.grant(_proverId, 1);
            emit DataProcessed(_projectId, _proverId, msg.sender, true, "");
        } catch Error(string memory revertReason) {
            emit DataProcessed(_projectId, _proverId, msg.sender, false, revertReason);
        }
    }

    function bindDapp(uint256 _projectId, address _dapp) external override onlyProjectOwner(_projectId) {
        dapp[_projectId] = _dapp;
        emit DappBound(_projectId, msg.sender, _dapp);
    }

    function unbindDapp(uint256 _projectId) external override onlyProjectOwner(_projectId) {
        delete dapp[_projectId];
        emit DappUnbound(_projectId, msg.sender);
    }
}
