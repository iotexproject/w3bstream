// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProjectStore.sol";
import "./interfaces/IProver.sol";
import "./interfaces/IRouter.sol";
import "./interfaces/IFleetManagement.sol";
import "./interfaces/IDapp.sol";

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract Router is IRouter, Initializable {
    address public override fleetManagement;
    address public override projectStore;

    mapping(uint256 => address) public override dapp;
    mapping(uint256 => uint256) public override credits;

    modifier onlyProjectOwner(uint256 _projectId) {
        require(
            IERC721(IFleetManagement(fleetManagement).project()).ownerOf(_projectId) == msg.sender,
            "not project owner"
        );
        _;
    }

    function initialize(address _fleetManagement, address _projectStore) public initializer {
        fleetManagement = _fleetManagement;
        projectStore = _projectStore;
    }

    function route(uint256 _projectId, uint256 _proverId, bytes calldata _data) external override {
        address _dapp = dapp[_projectId];
        require(_dapp != address(0), "no dapp");
        IFleetManagement _fm = IFleetManagement(fleetManagement);
        require(_fm.isActiveCoordinator(msg.sender, _projectId), "invalid coordinator");
        // TODO: validator prover signature
        // require(IProver(_fm.prover()).operator(_proverId) == msg.sender, "invalid prover operator");
        IProjectStore store = IProjectStore(projectStore);
        require(_projectId > 0 && _projectId <= store.count() && !store.isPaused(_projectId));
        require(_fm.isNormalProver(_proverId), "invalid prover");

        try IDapp(_dapp).process(_data) {
            credits[_proverId] += 1;
            emit DataProcessed(_projectId, _proverId, msg.sender, true, "");
        } catch Error(string memory revertReason) {
            emit DataProcessed(_projectId, _proverId, msg.sender, false, revertReason);
        }
    }

    function bindDapp(uint256 _projectId, address _dapp) external override onlyPorjectOwner(_projectId) {
        dapp[_projectId] = _dapp;
        emit DappBound(_projectId, msg.sender, _dapp);
    }

    function unbindDapp(uint256 _projectId) external override onlyProjectOwner(_projectId) {
        delete dapp[_projectId];
        emit DappUnbound(_projectId, msg.sender);
    }
}
