// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProjectRegistrar.sol";
import "./interfaces/IProject.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract ProjectRegistrar is IProjectRegistrar, Ownable {
    uint256 public override registrationFee;
    IProject public immutable project;

    constructor(address _project, uint256 _registrationFee) {
        project = IProject(_project);
        registrationFee = _registrationFee;
        emit RegistrationFeeSet(_registrationFee);
    }

    function register(string calldata _uri, bytes32 _hash) external payable override returns (uint256 _projectId) {
        require(msg.value >= registrationFee, "insufficient fee");
        _projectId = project.mint(msg.sender, _uri, _hash);

        emit ProjectRegister(_projectId);
    }

    function withdrawFee(address _account, uint256 _amount) external onlyOwner {
        (bool success, ) = payable(_account).call{value: _amount}("");
        require(success, "withdraw fee fail");

        emit WithdrawnFee(_account, _amount);
    }
}
