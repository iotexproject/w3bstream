// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract W3bstreamProver2 is OwnableUpgradeable {
    event BeneficiarySet(address indexed prover, address indexed beneficiary);
    event VMTypeAdded(address indexed prover, uint256 typ);
    event VMTypeDeleted(address indexed prover, uint256 typ);
    event ProverPaused(address indexed prover);
    event ProverResumed(address indexed prover);
    event RebateRatioSet(address indexed prover, uint16 ratio);

    mapping(address => mapping(uint256 => bool)) vmTypes;
    mapping(address => uint16) rebateRatios;
    mapping(address => address) beneficiaries;
    mapping(address => bool) paused;

    function initialize() public initializer {
        __Ownable_init();
    }

    function isVMTypeSupported(address _prover, uint256 _type) external view returns (bool) {
        return vmTypes[_prover][_type];
    }

    function beneficiary(address _prover) external view returns (address) {
        return beneficiaries[_prover];
    }

    function isPaused(address _prover) external view returns (bool) {
        return paused[_prover];
    }

    function rebateRatio(address _prover) external view returns (uint16) {
        return rebateRatios[_prover];
    }

    function register() external {
        address sender = msg.sender;
        require(beneficiaries[sender] == address(0), "already registered");
        beneficiaries[sender] = sender;
        emit BeneficiarySet(sender, sender);
    }

    function setRebateRatio(uint16 _ratio) external {
        address sender = msg.sender;
        rebateRatios[sender] = _ratio;
        emit RebateRatioSet(sender, _ratio);
    }

    function addVMType(uint256 _type) external {
        address sender = msg.sender;
        vmTypes[sender][_type] = true;
        emit VMTypeAdded(sender, _type);
    }

    function delVMType(uint256 _type) external {
        address sender = msg.sender;
        vmTypes[sender][_type] = false;
        emit VMTypeDeleted(sender, _type);
    }

    function changeBeneficiary(address _beneficiary) external {
        address sender = msg.sender;
        require(_beneficiary != address(0), "zero address");
        beneficiaries[sender] = _beneficiary;
        emit BeneficiarySet(sender, _beneficiary);
    }

    function pause() external {
        address sender = msg.sender;
        require(!paused[sender], "already paused");

        paused[sender] = true;
        emit ProverPaused(sender);
    }

    function resume() external {
        address sender = msg.sender;
        require(paused[sender], "already actived");

        paused[sender] = false;
        emit ProverResumed(sender);
    }
}
