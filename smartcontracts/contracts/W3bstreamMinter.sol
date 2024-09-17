// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

interface IDAO {
    function tip() external view returns (uint256, bytes32, uint256);
    function mint(bytes32 hash, uint256 timestamp) external;
}

interface ITaskManager {
    function assign(TaskAssignment[] calldata assignments, uint256 deadline) external;
}

struct BlockInfo {
    bytes4 meta;
    bytes32 prevhash;
    bytes32 merkleRoot;
    bytes4 difficulty;
    bytes8 nonce;
}

struct Sequencer {
    address addr;
    address operator;
    address beneficiary;
}

struct TaskAssignment {
    uint64 projectId;
    uint64 taskId;
    address prover;
}

contract W3bstreamMinter is OwnableUpgradeable {
    event TaskAllowanceSet(uint256 allowance);
    event TargetDurationSet(uint256 duration);
    event DifficultySet(bytes8 difficulty);

    IDAO public dao;
    ITaskManager public tm;
    uint256 public taskAllowance;
    uint256 public targetDuration;
    bytes4 public adhocDifficulty;
    bytes4 public currentDifficulty;

    uint32 private constant UPPER_BOUND = 0xffffffff;
    uint32 private constant LOWER_BOUND = 0x00000001;

    uint256[10] private durations;
    uint256 private durationSum;
    uint256 private durationNum;
    uint256 private durationIndex;

    function initialize(IDAO _dao, ITaskManager _tm) public initializer {
        __Ownable_init();
        dao = _dao;
        tm = _tm;
        _setTaskAllowance(720);
        _setTargetDuration(12);
        _setAdhocDifficulty(0x0fffffff);
    }

    function mint(
        BlockInfo calldata blockinfo, 
        Sequencer calldata coinbase,
        TaskAssignment[] calldata assignments
    ) public {
        require(coinbase.operator == msg.sender, "invalid operator");
        if (adhocDifficulty != 0) {
            require(blockinfo.difficulty == adhocDifficulty, "invalid difficulty");
        } else {
            require(blockinfo.difficulty == currentDifficulty, "invalid difficulty");
        }
        (, bytes32 tiphash, uint256 tipTimestamp) = dao.tip();
        require(tipTimestamp != block.number);
        require(blockinfo.prevhash == tiphash, "invalid prevhash");
        require(blockinfo.merkleRoot == keccak256(abi.encodePacked(coinbase.addr, coinbase.operator, coinbase.beneficiary)), "invalid merkle root");
        // TODO: review difficulty usage
        require(sha256(abi.encodePacked(blockinfo.meta, blockinfo.prevhash, blockinfo.merkleRoot, blockinfo.difficulty, blockinfo.nonce)) < blockinfo.difficulty, "invalid proof of work");
        bytes32 hash = keccak256(abi.encode(
            blockinfo.meta,
            blockinfo.prevhash,
            blockinfo.merkleRoot,
            blockinfo.difficulty,
            blockinfo.nonce,
            coinbase.addr,
            coinbase.operator,
            coinbase.beneficiary,
            assignments
        ));
        tm.assign(assignments, block.number + taskAllowance);
        _updateDifficulty(tipTimestamp);
        dao.mint(hash, block.number);
        // TODO: distribute block reward
    }

    function setTaskAllowance(uint256 allowance) public onlyOwner {
        _setTaskAllowance(allowance);
    }

    function _setTaskAllowance(uint256 allowance) internal {
        taskAllowance = allowance;
        emit TaskAllowanceSet(allowance);
    }

    function setTargetDuration(uint256 duration) public onlyOwner {
        _setTargetDuration(duration);
    }

    function _setTargetDuration(uint256 duration) internal {
        targetDuration = duration;
        emit TargetDurationSet(duration);
    }

    function setAdhocDifficulty(bytes4 difficulty) public onlyOwner {
        _setAdhocDifficulty(difficulty);
    }

    function _setAdhocDifficulty(bytes4 difficulty) internal {
        if (difficulty != 0) {
            _setDifficulty(difficulty);
        }
        adhocDifficulty = difficulty;
    }

    function _updateDifficulty(uint256 tipTimestamp) internal {
        uint256 duration = block.number - tipTimestamp;
        durationSum += duration - durations[durationIndex];
        durations[durationIndex] = duration;
        durationIndex = (durationIndex + 1) % durations.length;
        if (durationNum < durations.length) {
            durationNum++;
            return;
        }
        if (adhocDifficulty != 0) {
            return;
        }
        uint32 curr = uint32(currentDifficulty);
        uint40 next = curr;
        uint256 expectedSum = targetDuration * durationNum;
        if (durationSum * 5 > expectedSum * 6) {
            next *= 2;
        } else if (expectedSum * 4 > durationSum * 5) {
            next /= 2;
        } else {
            return;
        }
        if (next < LOWER_BOUND) {
            next = LOWER_BOUND;
        } else if (next > UPPER_BOUND) {
            next = UPPER_BOUND;
        }
        if (next != curr) {
            _setDifficulty(bytes4(uint32(next)));
        }
    }

    function _setDifficulty(bytes4 difficulty) internal {
        currentDifficulty = difficulty;
        emit DifficultySet(difficulty);
    }
}
