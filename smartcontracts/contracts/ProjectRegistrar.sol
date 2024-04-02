// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProjectRegistrar.sol";
import "./interfaces/IProjectStore.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract ProjectRegistrar is IProjectRegistrar, Ownable {
    uint256 public override registrationFee;
    IProjectStore public immutable projectStore;

    constructor(address _projectStore, uint256 _registrationFee) {
        projectStore = IProjectStore(_projectStore);
        registrationFee = _registrationFee;
        emit RegistrationFeeSet(_registrationFee);
    }

    function register(string calldata _uri, bytes32 _hash) external payable override returns (uint256 _projectId) {
        require(msg.value >= registrationFee, "insufficient fee");
        _projectId = projectStore.mint(msg.sender, _uri, _hash);

        emit ProjectRegister(_projectId);
    }

    function withdrawFee(address _account, uint256 _amount) external onlyOwner {
        (bool success, ) = payable(_account).call{value: _amount}("");
        require(success, "withdraw fee fail");

        emit WithdrawnFee(_account, _amount);
    }
}
