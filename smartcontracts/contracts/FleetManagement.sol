// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IProver.sol";
import "./interfaces/IFleetManagement.sol";

import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

contract FleetManagement is IFleetManagement, ReentrancyGuardUpgradeable, OwnableUpgradeable {
    uint256 public override epoch;
    uint256 public override minStake;
    address public override prover;
    address public override coordinator;
    address public override slasher;

    struct PendingWithdraw {
        uint256 timestamp;
        uint256 amount;
    }

    mapping(uint256 => uint256) public override stakedAmount;
    mapping(uint256 => PendingWithdraw) public pendingWithdraw;

    function initialize(uint256 _minStake) public initializer {
        __Ownable_init();
        __ReentrancyGuard_init();

        epoch = 1 hours;
        minStake = _minStake;
        prover = _prover;
    }

    function isActiveProver(uint256 _id) external view returns (bool) {
        IProver ip = IProver(prover);
        require(_id > 0 && _id < ip.count(), "invalid id");
        return !ip.isPaused(_id) && stakedAmount[_id] >= minStake;
    }

    function isActiveCoordinator(address _coordinator, uint256 _projectId) external view returns (bool) {
        return _coordinator == coordinator;
    }

    function setCoordinator(address _coordinator) external onlyOwner {
        coordinator = _coordinator;
        emit CoordinatorSet(_coordinator);
    }

    function setProver(address _prover) external onlyOwner {
        prover = _prover;
        emit ProverSet(_prover);
    }

    function setSlasher(address _slasher) external onlyOwner {
        slasher = _slasher;
        emit SlasherSet(_slasher);
    }

    function stake(uint256 _proverId) external payable override {
        require(msg.value > 0, "zero amount");
        require(IERC721(prover).ownerOf(_proverId) != address(0), "prove not exist");

        stakedAmount[_proverId] += msg.value;
        emit Stake(_proverId, msg.value);
    }

    function unstake(uint256 _proverId, uint256 _amount) external override {
        require(_amount > 0, "zero amount");
        require(pendingWithdraw[_proverId].timestamp == 0, "withdraw pending");
        require(IERC721(prover).ownerOf(_proverId) != msg.sender, "not owner");
        require(stakedAmount[_proverId] >= _amount, "invalid amount");

        stakedAmount[_proverId] -= _amount;
        PendingWithdraw storage _pending = pendingWithdraw[_proverId];
        _pending.timestamp = block.timestamp;
        _pending.amount = _amount;

        emit Unstake(_proverId, _amount);
    }

    function withdraw(uint256 _proverId, address _to) external override {
        require(IERC721(prover).ownerOf(_proverId) != msg.sender, "not owner");
        PendingWithdraw storage _pending = pendingWithdraw[_proverId];
        require(_pending.timestamp > 0 && _pending.timestamp + epoch <= block.timestamp, "invalid pending");

        uint256 _amount = _pending.amount;
        _pending.timestamp = 0;
        _pending.amount = 0;

        (bool success, ) = payable(_to).call{value: _amount}("");
        require(success, "withdraw fail");

        emit Withdrawn(_proverId, _to, _amount);
    }

    function grant(uint256 _proverId) external payable nonReentrant {
        // TODO grant token?
        address _owner = IERC721(prover).ownerOf(_proverId);
        require(_owner != address(0), "prover not exist");

        uint256 _amount = msg.value;
        (bool success, ) = payable(_owner).call{value: _amount}("");
        require(success, "withdraw fail");

        emit Grant(_proverId, _amount);
    }
}
