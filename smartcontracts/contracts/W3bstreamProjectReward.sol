// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IProject {
    function ownerOf(uint256 _projectId) external view returns (address);

    function isPaused(uint256 _projectId) external view returns (bool);
}

contract W3bstreamProjectReward is OwnableUpgradeable {
    event OperatorSet(address indexed operator);
    event RewardTokenSet(uint256 indexed id, address indexed rewardToken);
    event RewardAmountSet(address indexed owner, uint256 indexed id, uint256 amount);

    address public operator;
    address public project;

    mapping(uint256 => address) _rewardTokens;
    mapping(address => mapping(uint256 => uint256)) _rewardAmounts;

    modifier onlyOperator() {
        require(msg.sender == operator, "not operator");
        _;
    }

    function initialize(address _project) public initializer {
        __Ownable_init();
        project = _project;
    }

    function rewardToken(uint256 _id) external view returns (address) {
        return _rewardTokens[_id];
    }

    function rewardAmount(address owner, uint256 id) external view returns (uint256) {
        return _rewardAmounts[owner][id];
    }

    function isPaused(uint256 _projectId) external view returns (bool) {
        return IProject(project).isPaused(_projectId);
    }

    function setRewardToken(uint256 _id, address _rewardToken) external {
        require(IProject(project).ownerOf(_id) == msg.sender, "invalid project");
        require(_rewardToken != address(0), "zero address");
        require(_rewardTokens[_id] == address(0), "already set");
        _rewardTokens[_id] = _rewardToken;
        emit RewardTokenSet(_id, _rewardToken);
    }

    function setReward(uint256 _id, uint256 _amount) external {
        address sender = msg.sender;
        _rewardAmounts[sender][_id] = _amount;
        emit RewardAmountSet(sender, _id, _amount);
    }
}
