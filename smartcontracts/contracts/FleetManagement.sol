// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IMSP.sol";
import "./interfaces/IFleetManagement.sol";

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

interface IProverStore {
    function isPaused(uint256 _id) external view returns (bool);
    function prover(uint256 _id) external view returns (address);
    function mint(address _owner) external view returns (uint256);
}

interface ICreditCenter {
    function grant(address prover, uint256 amount) external;
}

contract FleetManagement is IFleetManagement, ReentrancyGuardUpgradeable, OwnableUpgradeable {
    uint256 public override epoch;
    uint256 public override minStake;
    uint256 public registrationFee;
    address public proverStore;
    address public coordinator;
    address public creditCenter;
    address public slasher;
    address public msp;

    function initialize(uint256 _minStake) public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();

        epoch = 1 hours;
        minStake = _minStake;
    }

    function setCreditCenter(address _creditCenter) external onlyOwner {
        creditCenter = _creditCenter;
        emit CreditCenterSet(_creditCenter);
    }

    function setCoordinator(address _coordinator) external onlyOwner {
        coordinator = _coordinator;
        emit CoordinatorSet(_coordinator);
    }

    function setProverStore(address _proverStore) external onlyOwner {
        proverStore = _proverStore;
        emit ProverStoreSet(_proverStore);
    }

    function setSlasher(address _slasher) external onlyOwner {
        slasher = _slasher;
        emit SlasherSet(_slasher);
    }

    function setMSP(address _msp) external onlyOwner {
        msp = _msp;
        emit MSPSet(_msp);
    }

    function setRegistrationFee(uint256 _fee) public onlyOwner {
        registrationFee = _fee;
        emit RegistrationFeeSet(_fee);
    }

    function withdrawFee(address payable _account, uint256 _amount) external onlyOwner {
        (bool success, ) = _account.call{value: _amount}("");
        require(success, "withdraw fee fail");

        emit FeeWithdrawn(_account, _amount);
    }

    function register() external payable nonReentrant returns (uint256) {
        require(msg.value >= registrationFee, "insufficient fee");
        return IProverStore(proverStore).mint(msg.sender);
    }

    function grant(uint256 _proverId, uint256 _amount) external override nonReentrant {
        address prover = IProverStore(proverStore).prover(_proverId);
        require(prover != address(0), "prover not exist");

        ICreditCenter(creditCenter).grant(prover, _amount);
    }

    function ownerOfProver(uint256 _proverId) external view override returns (address) {
        return IProverStore(proverStore).prover(_proverId);
    }

    function isActiveProver(uint256 _id) external view returns (bool) {
        IProverStore ps = IProverStore(proverStore);
        return !ps.isPaused(_id) && IMSP(msp).stakedAmount(ps.prover(_id)) >= minStake;
    }

    function isActiveCoordinator(address _coordinator, uint256 _projectId) external view returns (bool) {
        return _coordinator == coordinator;
    }
}
