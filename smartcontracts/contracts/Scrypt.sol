// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Scrypt {
    function hash(
        bytes calldata password,
        bytes calldata salt,
        uint64 N,
        uint32 r,
        uint32 p,
        uint32 keyLen,
        uint32 mode
    ) external view returns (bytes memory retval) {
        bytes memory input = abi.encode(password, salt, N, r, p, keyLen, mode);
        uint length = input.length;
        retval = new bytes(keyLen);
        assembly {
            // free memory pointer
            let success := staticcall(gas(), 0x8002, input, length, retval, keyLen)
            switch success
            case 0 {
                revert(0x0, 0x0)
            }
        }
    }
}