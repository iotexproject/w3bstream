// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

interface IDAO {
    function tip() external view returns (uint256, bytes32, uint256);
    function mint(bytes32 hash, uint256 timestamp) external;
}

interface ITaskManager {
    function assign(TaskAssignment[] calldata assignments, address sequencer, uint256 deadline) external;
}

interface IBlockRewardDistributor {
    function distribute(address recipient, uint256 amount) external;
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

contract W3bstreamBlockMinter is OwnableUpgradeable {
    event BlockRewardfSet(uint256 reward);
    event TaskAllowanceSet(uint256 allowance);
    event TargetDurationSet(uint256 duration);
    event NBitsSet(uint32 nbits);

    IDAO public dao;
    ITaskManager public taskManager;
    IBlockRewardDistributor public distributor;

    uint256 public blockReward;
    uint256 public taskAllowance;
    uint256 public targetDuration;
    bool public useAdhocNBits;
    uint32 public currentNBits;

    uint32 private constant MAX_EXPONENT = 0x1c;
    uint32 private constant UPPER_BOUND = 0xffff00;
    uint32 private constant LOWER_BOUND = 0x8000;
    uint256 private constant MAX_TARGET = 0x00000000ffff0000000000000000000000000000000000000000000000000000;
    
    uint256 private _currentTarget;
    uint256[10] private _durations;
    uint256 private _durationSum;
    uint256 private _durationNum;
    uint256 private _durationIndex;

    function initialize(IDAO _dao, ITaskManager _taskManager, IBlockRewardDistributor _distributor) public initializer {
        __Ownable_init();
        dao = _dao;
        taskManager = _taskManager;
        distributor = _distributor;
        _setBlockReward(1000000000000000000);
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
        require(target == _currentTarget, "invalid nbits");
        (, bytes32 tiphash, uint256 tipTimestamp) = dao.tip();
        require(tipTimestamp != block.number);
        require(blockinfo.prevhash == tiphash, "invalid prevhash");
        require(blockinfo.merkleRoot == keccak256(abi.encodePacked(coinbase.addr, coinbase.operator, coinbase.beneficiary)), "invalid merkle root");
        // TODO: review target usage
        bytes memory header = abi.encodePacked(blockinfo.meta, blockinfo.prevhash, blockinfo.merkleRoot, blockinfo.nbits, blockinfo.nonce);
        require(uint256(sha256(header)) < target, "invalid proof of work");
        bytes32 h = keccak256(abi.encode(header, assignments));
        taskManager.assign(assignments, coinbase.beneficiary, block.number + taskAllowance);
        _updateTarget(tipTimestamp);
        dao.mint(h, block.number);
        distributor.distribute(coinbase.beneficiary, blockReward);
    }
/*
    function scrypt(bytes calldata x) public pure returns (bytes32) {
        require(x.length == 80, "invalid length");
        uint64 keyLen = 32;
        bytes memory b = pbkdf2(x, x, 128);
        // bytes memory b = pbkdf2_Key(x, x, 1, 128);
        smix(b);
        uint32 value = uint32(uint8(b[0])) << 24 | uint32(uint8(b[1])) << 16 | uint32(uint8(b[2])) << 8 | uint32(uint8(b[3]));
        uint32 addend = 0xe0;
        uint32 newValue = value + addend;
        uint32 high17 = newValue >> 15;
        uint32 low15 = newValue & 0x7fff;
        uint32 result = (high17 << 15) | low15;
        b[0] = bytes1(uint8(result >> 24));
        b[1] = bytes1(uint8(result >> 16));
        b[2] = bytes1(uint8(result >> 8));
        b[3] = bytes1(uint8(result));

        // return reverseArray(pbkdf2_Key(x, b, 1, keyLen));
        bytes memory reverse = reverseArray(pbkdf2(x, b, keyLen));
        bytes32 retval;
        assembly {
            retval := mload(add(reverse, 32))
        }
        return retval;
    }
*/
    function reverseArray(bytes memory b) public pure returns (bytes memory) {
        bytes1 tmp;
        for (uint i = 0; i < b.length / 2; i++) {
            tmp = b[i];
            b[i] = b[b.length - i - 1];
            b[b.length - i - 1] = tmp;
        }
        return b;
    }
/*
    function pbkdf2_Key(bytes calldata x, bytes calldata salt, uint64 iter, uint64 keyLen) public pure returns (bytes memory) {
        bytes32 h = sha256(x);
        uint32 hashLen = 32;
        uint64 numBlocks = 4;
        bytes memory buf = new bytes(4);
        bytes memory dk = new bytes(128);
        for (uint i = 1; i <= numBlocks; i++) {
            h = 0x6A09E667BB67AE853C6EF372A54FF53A510E527F9B05688C1F83D9AB5BE0CD19;
            // h.Write(salt);
            buf[0] = bytes1(uint8(i >> 24));
            buf[1] = bytes1(uint8(i >> 16));
            buf[2] = bytes1(uint8(i >> 8));
            buf[3] = bytes1(uint8(i));
            // h.Write(buf[:4]);
            // dk = h.Sum(dk);
            for (uint j = 0; j < iter; j++) {
                // h.Reset();
                // h.Write(dk);
                // dk = h.Sum(dk);
            }
            for (uint j = 0; j < hashLen; j++) {
                dk[i * hashLen + j] = dk[j];
            }
            // buf = sha256(abi.encodePacked(h, y));
            // dk[i] = buf;
        }
        bytes memory retval = new bytes(keyLen);
        for (uint i = 0; i < keyLen; i++) {
            retval[i] = dk[i];
        }
        return retval;
    }
*/
    function setBlockReward(uint256 reward) public onlyOwner {
        _setBlockReward(reward);
    }

    function _setBlockReward(uint256 reward) internal {
        blockReward = reward;
        emit BlockRewardfSet(reward);
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

/*
    function blockXOR(uint32[] memory dst, uint32[] memory src, uint256 offset, uint64 len) pure internal returns (uint32[] memory) {
        for (uint i = 0; i < len; i++) {
            dst[i] ^= src[i + offset];
        }
        return dst;
    }

    function integer(uint32[] memory b) pure internal returns (uint64) {
        return uint64(b[16]) | uint64(b[17]) << 32;
    }

    function salsaXOR(uint32[16] memory tmp, uint32[] memory input, uint offset) pure internal returns (uint32[16] memory output) {
        uint32[16] memory z;
        uint32[16] memory x;
        for (uint i = 0; i < 16; i++) {
            z[i] = tmp[i] ^ input[offset + i];
            x[i] = z[i];
        }
        for (uint i = 0; i < 8; i += 2) {
            unchecked {
            uint32 u = x[0] + x[12];
            x[4] ^= u << 7 | u >> 25;
            u = x[4] + x[0];
            x[8] ^= u << 9 | u >> 23;
            u = x[8] + x[4];
            x[12] ^= u << 13 | u >> 19;
            u = x[12] + x[8];
            x[0] ^= u << 18 | u >> 14;

            u = x[5] + x[1];
            x[9] ^= u << 7 | u >> 25;
            u = x[9] + x[5];
            x[13] ^= u << 9 | u >> 23;
            u = x[13] + x[9];
            x[1] ^= u << 13 | u >> 19;
            u = x[1] + x[13];
            x[5] ^= u << 18 | u >> 14;

            u = x[10] + x[6];
            x[14] ^= u << 7 | u >> 25;
            u = x[14] + x[10];
            x[2] ^= u << 9 | u >> 23;
            u = x[2] + x[14];
            x[6] ^= u << 13 | u >> 19;
            u = x[6] + x[2];
            x[10] ^= u << 18 | u >> 14;

            u = x[15] + x[11];
            x[3] ^= u << 7 | u >> 25;
            u = x[3] + x[15];
            x[7] ^= u << 9 | u >> 23;
            u = x[7] + x[3];
            x[11] ^= u << 13 | u >> 19;
            u = x[11] + x[7];
            x[15] ^= u << 18 | u >> 14;

            u = x[0] + x[3];
            x[1] ^= u << 7 | u >> 25;
            u = x[1] + x[0];
            x[2] ^= u << 9 | u >> 23;
            u = x[2] + x[1];
            x[3] ^= u << 13 | u >> 19;
            u = x[3] + x[2];
            x[0] ^= u << 18 | u >> 14;

            u = x[5] + x[4];
            x[6] ^= u << 7 | u >> 25;
            u = x[6] + x[5];
            x[7] ^= u << 9 | u >> 23;
            u = x[7] + x[6];
            x[4] ^= u << 13 | u >> 19;
            u = x[4] + x[7];
            x[5] ^= u << 18 | u >> 14;

            u = x[10] + x[9];
            x[11] ^= u << 7 | u >> 25;
            u = x[11] + x[10];
            x[8] ^= u << 9 | u >> 23;
            u = x[8] + x[11];
            x[9] ^= u << 13 | u >> 19;
            u = x[9] + x[8];
            x[10] ^= u << 18 | u >> 14;

            u = x[15] + x[14];
            x[12] ^= u << 7 | u >> 25;
            u = x[12] + x[15];
            x[13] ^= u << 9 | u >> 23;
            u = x[13] + x[12];
            x[14] ^= u << 13 | u >> 19;
            u = x[14] + x[13];
            x[15] ^= u << 18 | u >> 14;
            }
        }
        unchecked {
        for (uint i = 0; i < 16; i++) {
            x[i] += z[i];
        }
        }
        for (uint i = 0; i < 16; i++) {
            output[i] = x[i];
        }
    }

    function blockMix(uint32[16] memory tmp, uint32[] memory x) pure internal returns (uint32[16] memory, uint32[] memory) {
        for (uint i = 0; i < 16; i++) {
            tmp[i] = x[16 + i];
        }
        uint32[] memory output = new uint32[](32);
        uint32[16] memory v = salsaXOR(tmp, x, 0);
        for (uint i = 0; i < 16; i++) {
            tmp[i] = output[i];
            output[i] = v[i];
        }

        v = salsaXOR(tmp, x, 16);
        for (uint i = 0; i < 16; i++) {
            tmp[i] = output[16 + i];
            output[16 + i] = v[i];
        }

        return (tmp, output);
    }

    function smix(bytes memory b) pure internal returns (bytes memory) {
        uint32[16] memory tmp;
        uint32[] memory x = new uint32[](32);
        uint32[] memory y = new uint32[](32);
        uint32[] memory v = new uint32[](32 * 1024);

        uint j = 0;
        for (uint i = 0; i < 32; i++) {
            x[i] = uint32(uint8(b[j])) | uint32(uint8(b[j+1]))<<8 | uint32(uint8(b[j+2]))<<16 | uint32(uint8(b[j+3]))<<24;
            j += 4;
        }
        for (uint i = 0; i < 1024; i += 2) {
            for (uint k = 0; k < 32; k++) {
                v[i * 32 + k] = x[k];
            }
            (tmp, y) = blockMix(tmp, x);
            for (uint k = 0; k < 32; k++) {
                v[(i + 1) * 32 + k] = y[k];
            }
            (tmp, x) = blockMix(tmp, y);
        }
        for (uint i = 0; i < 1024; i += 2) {
            j = uint32(integer(x) & uint64(1023));
            blockXOR(x, v, j * 32, 32);
            (tmp, y) = blockMix(tmp, x);
            j = uint32(integer(y) & uint64(1023));
            blockXOR(y, v, j * 32, 32);
            (tmp, x) = blockMix(tmp, y);
        }
        j = 0;
        for (uint i = 0; i < 32; i++) {
            b[j + 0] = bytes1(uint8(x[i] >> 0));
            b[j + 1] = bytes1(uint8(x[i] >> 8));
            b[j + 2] = bytes1(uint8(x[i] >> 16));
            b[j + 3] = bytes1(uint8(x[i] >> 24));
            j += 4;
        }

        return b;
    }

    function hmacsha256(bytes calldata key, bytes memory message) pure public returns (bytes32) {
        bytes32 keyl;
        bytes32 keyr;
        uint i;
        if (key.length > 64) {
            keyl = sha256(key);
        } else {
            for (i = 0; i < key.length && i < 32; i++)
                keyl |= bytes32(uint256(uint8(key[i])) * 2**(8 * (31 - i)));
            for (i = 32; i < key.length && i < 64; i++)
                keyr |= bytes32(uint256(uint8(key[i])) * 2**(8 * (63 - i)));
        }
        bytes32 threesix = 0x3636363636363636363636363636363636363636363636363636363636363636;
        bytes32 fivec = 0x5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c5c; 
        return sha256(abi.encode(fivec ^ keyl, fivec ^ keyr, sha256(abi.encode(threesix ^ keyl, threesix ^ keyr, message))));
    }

    function pbkdf2(bytes calldata k, bytes memory salt, uint dklen) public pure returns (bytes memory) {
        bytes memory m = new bytes(salt.length + 4);
        for (uint i = 0; i < salt.length; i++) {
            m[i] = salt[i];
        }
        bytes32[4] memory r;
        for (uint i = 0; i * 32 < dklen; i++) {
            m[m.length - 1] = bytes1(uint8(i + 1));
            r[i] = hmacsha256(k, m);
        }
        bytes memory retval = new bytes(dklen);
        for (uint i = 0; i < dklen; i++) {
            retval[i] = r[i / 32][i % 32];
        }
        return retval;
    }
*/
}
