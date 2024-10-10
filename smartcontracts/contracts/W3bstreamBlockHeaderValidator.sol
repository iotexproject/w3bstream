// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {BlockHeader, IBlockHeaderValidator, IScrypt} from "./interfaces/IBlockHeaderValidator.sol";

contract W3bstreamBlockHeaderValidator is IBlockHeaderValidator, Ownable {
    event OperatorSet(address operator);
    event TargetDurationSet(uint256 duration);
    event NBitsSet(uint32 nbits);

    uint32 public constant MAX_EXPONENT = 0x1c;
    uint32 public constant UPPER_BOUND = 0xffff00;
    uint32 public constant LOWER_BOUND = 0x8000;
    // uint256 public constant MAX_TARGET = 0x00000000ffff0000000000000000000000000000000000000000000000000000;

    address public operator;
    IScrypt public scrypt;
    uint256 public targetDuration;
    uint32 public currentNBits;
    bool public useAdhocNBits;

    uint256 private _currentTarget;
    uint256[10] private _durations;
    uint256 private _durationSum;
    uint256 private _durationNum;
    uint256 private _durationIndex;

    constructor(IScrypt _scrypt) Ownable() {
        scrypt = _scrypt;
        _setTargetDuration(12);
        _setAdhocNBits(0x1c7fffff);
    }

    function validate(BlockHeader calldata header) public view returns (bytes memory) {
        require(header.nbits == currentNBits, "invalid nbits");
        bytes memory encodedHeader = abi.encodePacked(header.meta, header.prevhash, header.merkleRoot, header.nbits, header.nonce);
        bytes memory headerHash = scrypt.hash(encodedHeader, encodedHeader, 1024, 1, 1, 32, 224);
        require(headerHash.length == 32, "invalid header hash length");
        require(uint256(bytes32(headerHash)) <= _currentTarget, "invalid proof of work");
        return encodedHeader;
    }

    function setOperator(address _operator) public onlyOwner {
        operator = _operator;
        emit OperatorSet(_operator);
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

    function updateDuration(uint256 duration) public {
        require(msg.sender == operator, "not operator");
        _durationSum += duration - _durations[_durationIndex];
        _durations[_durationIndex] = duration;
        _durationIndex = (_durationIndex + 1) % _durations.length;
        if (_durationNum < _durations.length) {
            _durationNum++;
            return;
        }
        if (useAdhocNBits) {
            return;
        }
        uint32 nbits = uint32(currentNBits);
        uint32 next = _nextNBits(nbits, targetDuration * _durationNum, _durationSum);
        if (next != nbits) {
            _setNBits(next);
        }
    }

    function _nextNBits(uint32 nbits, uint256 expectedSum, uint256 sum) internal pure returns (uint32) {
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
        _currentTarget = target;
        emit NBitsSet(nbits);
    }

    function decodeNBits(uint32 nbits) internal pure returns (uint32, uint32) {
        return (nbits >> 24, nbits & 0x00ffffff);
    }

    function nbitsToTarget(uint32 nbits) public pure returns (uint256) {
        (uint32 exponent, uint256 coefficient) = decodeNBits(nbits);
        require(exponent <= MAX_EXPONENT, "invalid nbits");
        require(coefficient >= LOWER_BOUND && coefficient <= UPPER_BOUND, "invalid nbits");
        return coefficient << (8 * (exponent - 3));
    }
}
