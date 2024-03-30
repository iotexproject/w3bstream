// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProject.sol";
import "./interfaces/IProver.sol";
import "./interfaces/IRouter.sol";
import "./interfaces/IFleetManagement.sol";
import "./interfaces/IDapp.sol";

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract Router is IRouter, Initializable {
    address public override fleetManagement;
    mapping(uint256 => address) public override dapp;
    mapping(uint256 => uint256) public override credits;

    function initialize(address _fleetManagement) public initializer {
        fleetManagement = _fleetManagement;
    }

    function route(uint256 _projectId, uint256 _proverId, bytes calldata _data) external override {
        address _dapp = dapp[_projectId];
        require(_dapp != address(0), "no dapp");
        IFleetManagement _fm = IFleetManagement(fleetManagement);
        require(IProver(_fm.prover()).operator(_proverId) == msg.sender, "invalid prover operator");
        require(_fm.isNormalProject(_projectId), "invalid project");
        require(_fm.isNormalProver(_proverId), "invalid prover");

        try IDapp(_dapp).process(_data) {
            credits[_proverId] += 1;
            emit DataProcessed(_projectId, _proverId, msg.sender, true, "");
        } catch Error(string memory revertReason) {
            emit DataProcessed(_projectId, _proverId, msg.sender, false, revertReason);
        }
    }

    function bindDapp(uint256 _projectId, address _dapp) external override {
        require(
            IProject(IFleetManagement(fleetManagement).project()).operators(_projectId, msg.sender),
            "not operator"
        );

        dapp[_projectId] = _dapp;
        emit BindDapp(_projectId, msg.sender, _dapp);
    }

    function unbindDapp(uint256 _projectId) external override {
        require(
            IProject(IFleetManagement(fleetManagement).project()).operators(_projectId, msg.sender),
            "not operator"
        );

        delete dapp[_projectId];
        emit UnbindDapp(_projectId, msg.sender);
    }
}
