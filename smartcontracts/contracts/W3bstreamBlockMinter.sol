// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {BlockHeader, IBlockHeaderValidator} from "./interfaces/IBlockHeaderValidator.sol";
import {ITaskManager, TaskAssignment} from "./interfaces/ITaskManager.sol";

interface IDAO {
    function tip() external view returns (uint256, bytes32, uint256);

    function mint(bytes32 hash, uint256 timestamp) external;
}

interface IBlockRewardDistributor {
    function distribute(address recipient, uint256 amount) external;
}

struct Sequencer {
    address addr;
    address operator;
    address beneficiary;
}

contract W3bstreamBlockMinter is OwnableUpgradeable {
    event BlockRewardSet(uint256 reward);
    event TaskAllowanceSet(uint256 allowance);
    event TargetDurationSet(uint256 duration);
    event NBitsSet(uint32 nbits);

    IDAO public dao;
    ITaskManager public taskManager;
    IBlockRewardDistributor public distributor;
    IBlockHeaderValidator public headerValidator;

    uint256 public blockReward;
    uint256 public taskAllowance;

    function initialize(
        IDAO _dao,
        ITaskManager _taskManager,
        IBlockRewardDistributor _distributor,
        IBlockHeaderValidator _headerValidator
    ) public initializer {
        __Ownable_init();
        dao = _dao;
        taskManager = _taskManager;
        distributor = _distributor;
        headerValidator = _headerValidator;
        _setBlockReward(1000000000000000000);
        _setTaskAllowance(720);
    }

    function mint(
        BlockHeader calldata header,
        Sequencer calldata coinbase,
        TaskAssignment[] calldata assignments
    ) public {
        //require(false, "i am here 1");
        require(coinbase.operator == msg.sender, "invalid operator");
        (uint256 tipBlockNumber, bytes32 tiphash, ) = dao.tip();
        require(tipBlockNumber != block.number);
        require(header.prevhash == tiphash, "invalid prevhash");
        require(
            header.merkleRoot == keccak256(abi.encode(coinbase.addr, coinbase.operator, coinbase.beneficiary)),
            "invalid merkle root"
        );
        bytes memory encodedHeader = headerValidator.validate(header);
        //require(1 < 0, "i am here 2");
        bytes32 blockHash = keccak256(abi.encode(encodedHeader, assignments));
        taskManager.assign(assignments, coinbase.beneficiary, block.number + taskAllowance);
        headerValidator.updateDuration(block.number - tipBlockNumber);
        dao.mint(blockHash, block.number);
        distributor.distribute(coinbase.beneficiary, blockReward);
    }

    function setBlockReward(uint256 reward) public onlyOwner {
        _setBlockReward(reward);
    }

    function _setBlockReward(uint256 reward) internal {
        blockReward = reward;
        emit BlockRewardSet(reward);
    }

    function setTaskAllowance(uint256 allowance) public onlyOwner {
        _setTaskAllowance(allowance);
    }

    function _setTaskAllowance(uint256 allowance) internal {
        taskAllowance = allowance;
        emit TaskAllowanceSet(allowance);
    }
}
