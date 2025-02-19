// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IRouter.sol";
import "./interfaces/IDapp.sol";
import "./interfaces/IIoIDProxy.sol";

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {ITaskManager} from "./interfaces/ITaskManager.sol";

interface IProjectStore {
    function isPaused(uint256 _projectId) external view returns (bool);
}

interface IProverStore {
    function isPaused(address prover) external view returns (bool);
}

contract W3bstreamRouter is IRouter, Initializable {
    ITaskManager public taskManager;
    IProverStore public proverStore;
    address public projectStore;

    mapping(uint256 => address) public override dapp;

    modifier onlyProjectOwner(uint256 _projectId) {
        address projectOwner = IERC721(projectStore).ownerOf(_projectId);
        require(projectOwner == msg.sender || IIoIDProxyOwner(projectOwner).owner() == msg.sender, "not project owner");
        _;
    }

    function initialize(
        ITaskManager _taskManager,
        IProverStore _proverStore,
        address _projectStore
    ) public initializer {
        taskManager = _taskManager;
        projectStore = _projectStore;
        proverStore = _proverStore;
    }

    function route(
        address _prover,
        uint256 _projectId,
        bytes32[] calldata _taskIds,
        bytes calldata _data
    ) external override {
        address _dapp = dapp[_projectId];
        require(_dapp != address(0), "no dapp");
        require(!proverStore.isPaused(_prover), "prover paused");
        require(!IProjectStore(projectStore).isPaused(_projectId), "project paused");

        IDapp(_dapp).process(_prover, _projectId, _taskIds, _data);
        for (uint256 i = 0; i < _taskIds.length; i++) {
            taskManager.settle(_projectId, _taskIds[i], _prover);
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
