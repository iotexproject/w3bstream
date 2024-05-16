// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

interface IProjectStore {
    function bind(uint256 _projectId) external;
}

contract ProjectRegistrar is OwnableUpgradeable {
    event ProjectRegistered(uint256 indexed projectId);
    event RegistrationFeeSet(uint256 fee);
    event FeeWithdrawn(address indexed account, uint256 amount);

    uint256 public registrationFee;
    IProjectStore public projectStore;

    function initialize(address _projectStore) public initializer {
        __Ownable_init();
        projectStore = IProjectStore(_projectStore);
    }

    function setRegistrationFee(uint256 _fee) public onlyOwner {
        registrationFee = _fee;
        emit RegistrationFeeSet(_fee);
    }

    function register(uint256 _projectId) external payable {
        require(msg.value >= registrationFee, "insufficient fee");
        projectStore.bind(_projectId);
    }

    function withdrawFee(address payable _account, uint256 _amount) external onlyOwner {
        (bool success, ) = _account.call{value: _amount}("");
        require(success, "withdraw fee fail");

        emit FeeWithdrawn(_account, _amount);
    }
}
