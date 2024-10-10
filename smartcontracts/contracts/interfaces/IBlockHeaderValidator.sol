// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IBlockHeaderValidator {
    function validate(BlockHeader calldata header) external view returns (bytes memory);
    function updateDuration(uint256 duration) external;
}

interface IScrypt {
    function hash(
        bytes calldata password,
        bytes calldata salt,
        uint64 N,
        uint32 r,
        uint32 p,
        uint32 keyLen,
        uint32 mode
    ) external view returns (bytes memory);
}

struct BlockHeader {
    bytes4 meta;
    bytes32 prevhash;
    bytes32 merkleRoot;
    uint32 nbits;
    bytes8 nonce;
}
