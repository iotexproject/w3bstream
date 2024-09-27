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
        assembly {
            if iszero(staticcall(gas(), 0x8002, add(input, 0x20), length, 0, 0)) {
                revert(0x0, 0x0)
            }
            if iszero(eq(returndatasize(), keyLen)) {
                revert(0x0, 0x0)
            }
            mstore(retval, returndatasize()) // Store the length.
            let o := add(retval, 0x20)
            returndatacopy(o, 0x00, returndatasize()) // Copy the returndata.
            mstore(0x40, add(o, returndatasize())) // Allocate the memory.
        }

        return retval;
    }
}