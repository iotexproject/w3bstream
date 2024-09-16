// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

interface IDAO {
    function tip() external view returns (uint256, bytes32, uint256);
    function mint(bytes32 hash, uint256 timestamp) external;
}

interface ITaskManager {
    function assign(uint64 projectId, uint64 taskId, address prover, uint256 deadline) external;
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
    IDAO public dao;
    ITaskManager public tm;

    function initialize(IDAO _dao, ITaskManager _tm) public initializer {
        __Ownable_init();
        dao = _dao;
        tm = _tm;
    }

    function mint(
        BlockInfo calldata blockinfo, 
        uint256 timestamp,
        Sequencer calldata coinbase,
        TaskAssignment[] calldata assignments
    ) public {
        require(coinbase.operator == msg.sender, "invalid operator");
        (, bytes32 tip, ) = dao.tip();
        require(blockinfo.prevhash == tip, "invalid prevhash");
        // TODO: timestamp is not larger than block.timestamp
        require(timestamp > block.timestamp - 1 minutes, "invalid timestamp");
        require(blockinfo.merkleRoot == keccak256(abi.encodePacked(timestamp, coinbase.addr, coinbase.operator, coinbase.beneficiary)), "invalid merkle root");
        // TODO: review difficulty usage
        require(sha256(abi.encodePacked(blockinfo.meta, blockinfo.prevhash, blockinfo.merkleRoot, blockinfo.difficulty, blockinfo.nonce)) < blockinfo.difficulty, "invalid proof of work");
        bytes32 hash = keccak256(abi.encode(
            blockinfo.meta,
            blockinfo.prevhash,
            blockinfo.merkleRoot,
            blockinfo.difficulty,
            blockinfo.nonce,
            timestamp,
            coinbase.addr,
            coinbase.operator,
            coinbase.beneficiary,
            assignments
        ));
        uint256 deadline = block.timestamp + 1 hours;
        for (uint i = 0; i < assignments.length; i++) {
            tm.assign(assignments[i].projectId, assignments[i].taskId, assignments[i].prover, deadline);
        }
        // TODO: adjust difficulty
        dao.mint(hash, timestamp);
        // TODO: distribute block reward
    }
}
