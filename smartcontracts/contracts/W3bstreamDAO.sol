// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract W3bstreamDAO is OwnableUpgradeable {
    struct Block {
        bytes32 hash;
        uint256 timestamp;
    }
    event BlockAdded(uint256 indexed num, bytes32 hash, uint256 timestamp);

    Block[] public blocks;

    function initialize(bytes32 genesis) public initializer {
        __Ownable_init();
        _mint(genesis, block.timestamp);
    }

    function mint(bytes32 hash, uint256 timestamp) public onlyOwner {
        _mint(hash, timestamp);
    }

    function tip() public view returns (uint256, bytes32, uint256) {
        uint256 blocknum = blocks.length - 1;
        Block storage tipblock = blocks[blocknum];
        return (blocknum, tipblock.hash, tipblock.timestamp);
    }

    function _mint(bytes32 hash, uint256 timestamp) internal {
        blocks.push(Block({timestamp: timestamp, hash: hash}));
        emit BlockAdded(blocks.length - 1, hash, timestamp);
    }
}
