// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ITaskManager, TaskAssignment} from "./interfaces/ITaskManager.sol";

interface IRewardDistributor {
    function distribute(address recipient, uint256 amount) external;
}

struct Sequencer {
    address addr;
    address operator;
    address beneficiary;
}

contract W3bstreamMinter is OwnableUpgradeable {
    event RewardSet(uint256 reward);
    event TaskAllowanceSet(uint256 allowance);
    event TargetDurationSet(uint256 duration);

    ITaskManager public taskManager;
    IRewardDistributor public distributor;

    uint256 public reward;
    uint256 public taskAllowance;

    function initialize(ITaskManager _taskManager, IRewardDistributor _distributor) public initializer {
        __Ownable_init();
        taskManager = _taskManager;
        distributor = _distributor;
        _setReward(1000000000000000000);
        _setTaskAllowance(720);
    }

    function mint(Sequencer calldata coinbase, TaskAssignment[] calldata assignments) public {
        require(coinbase.operator == msg.sender, "invalid coinbase operator");
        taskManager.assign(assignments, coinbase.beneficiary, block.number + taskAllowance);
        distributor.distribute(coinbase.beneficiary, reward);
    }

    function setReward(uint256 _reward) public onlyOwner {
        _setReward(_reward);
    }

    function _setReward(uint256 _reward) internal {
        reward = _reward;
        emit RewardSet(reward);
    }

    function setTaskAllowance(uint256 allowance) public onlyOwner {
        _setTaskAllowance(allowance);
    }

    function _setTaskAllowance(uint256 allowance) internal {
        taskAllowance = allowance;
        emit TaskAllowanceSet(allowance);
    }
}
