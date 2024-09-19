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
    uint32 nbits;
    bytes8 nonce;
}

struct Sequencer {
    address addr;
    address operator;
    address beneficiary;
}

struct TaskAssignment {
    uint256 projectId;
    bytes32 taskId;
    address prover;
}

contract W3bstreamMinter is OwnableUpgradeable {
    event TaskAllowanceSet(uint256 allowance);
    event TargetDurationSet(uint256 duration);
    event NBitsSet(uint32 nbits);

    IDAO public dao;
    ITaskManager public tm;
    uint256 public taskAllowance;
    uint256 public targetDuration;
    bool public useAdhocNBits;
    uint32 public currentNBits;
    uint256 public currentTarget;

    uint32 private constant MAX_EXPONENT = 0x1c;
    uint32 private constant UPPER_BOUND = 0xffff00;
    uint32 private constant LOWER_BOUND = 0x8000;
    uint256 private constant MAX_TARGET = 0x00000000ffff0000000000000000000000000000000000000000000000000000;
    
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
        _setAdhocNBits(0x0f7fffff);
    }

    function mint(
        BlockInfo calldata blockinfo, 
        Sequencer calldata coinbase,
        TaskAssignment[] calldata assignments
    ) public {
        require(coinbase.operator == msg.sender, "invalid operator");
        uint256 target = nbitsToTarget(blockinfo.nbits);
        require(target == currentTarget, "invalid nbits");
        (, bytes32 tiphash, uint256 tipTimestamp) = dao.tip();
        require(tipTimestamp != block.number);
        require(blockinfo.prevhash == tiphash, "invalid prevhash");
        require(blockinfo.merkleRoot == keccak256(abi.encodePacked(coinbase.addr, coinbase.operator, coinbase.beneficiary)), "invalid merkle root");
        // TODO: review target usage
        bytes memory header = abi.encodePacked(blockinfo.meta, blockinfo.prevhash, blockinfo.merkleRoot, blockinfo.nbits, blockinfo.nonce);
        require(uint256(sha256(header)) < target, "invalid proof of work");
        bytes32 h = keccak256(abi.encode(header, assignments));
        tm.assign(assignments, block.number + taskAllowance);
        _updateTarget(tipTimestamp);
        dao.mint(h, block.number);
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

    function setAdhocNBits(uint32 nbits) public onlyOwner {
        _setAdhocNBits(nbits);
    }

    function _setAdhocNBits(uint32 nbits) internal {
        if (nbits == 0) {
            useAdhocNBits = false;
            return;
        }
        _setNBits(nbits);
        useAdhocNBits = true;
    }

    function _updateTarget(uint256 tipTimestamp) internal {
        uint256 duration = block.number - tipTimestamp;
        durationSum += duration - durations[durationIndex];
        durations[durationIndex] = duration;
        durationIndex = (durationIndex + 1) % durations.length;
        if (durationNum < durations.length) {
            durationNum++;
            return;
        }
        if (useAdhocNBits) {
            return;
        }
        uint32 nbits = uint32(currentNBits);
        uint32 next = nextNBits(nbits, targetDuration * durationNum, durationSum);
        if (next != nbits) {
            _setNBits(next);
        }
    }

    function nextNBits(uint32 nbits, uint256 expectedSum, uint256 sum) internal pure returns (uint32) {
        if (sum * 5 > expectedSum * 6) {
            (uint32 exponent, uint32 coefficient) = decodeNBits(nbits);
            if (coefficient < UPPER_BOUND) {
                return (exponent << 24) | uint32(coefficient + 1);
            }
            if (exponent < MAX_EXPONENT) {
                return ((exponent + 1) << 24) | LOWER_BOUND;
            }
        } else if (expectedSum * 4 > sum * 5) {
            (uint32 exponent, uint32 coefficient) = decodeNBits(nbits);
            if (coefficient > LOWER_BOUND) {
                return (exponent << 24) | uint32(coefficient - 1);
            }
            if (exponent > 0) {
                return ((exponent - 1) << 24) | UPPER_BOUND;
            }
        }
        return nbits;
    }

    function _setNBits(uint32 nbits) internal {
        uint256 target = nbitsToTarget(nbits);
        currentNBits = nbits;
        currentTarget = target;
        emit NBitsSet(nbits);
    }

    function decodeNBits(uint32 nbits) internal pure returns (uint32, uint32) {
        return (nbits >> 24, nbits & 0x00ffffff);
    }

    function nbitsToTarget(uint32 nbits) internal pure returns (uint256) {
        (uint32 exponent, uint256 coefficient) = decodeNBits(nbits);
        require(exponent <= MAX_EXPONENT, "invalid nbits");
        require(coefficient >= LOWER_BOUND && coefficient <= UPPER_BOUND, "invalid nbits");
        return coefficient << (8 * (exponent - 3));
    }
}
