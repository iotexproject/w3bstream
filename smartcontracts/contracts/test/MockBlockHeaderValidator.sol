// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {BlockHeader, IBlockHeaderValidator, IScrypt} from "../interfaces/IBlockHeaderValidator.sol";

contract MockBlockHeaderValidator is IBlockHeaderValidator, Ownable {
    event NBitsSet(uint32 nbits);
    event OperatorSet(address operator);

    address public operator;

    constructor() Ownable() {}

    function validate(BlockHeader calldata header) public pure returns (bytes memory) {
        bytes memory encodedHeader = abi.encodePacked(
            header.meta,
            header.prevhash,
            header.merkleRoot,
            header.nbits,
            header.nonce
        );
        return encodedHeader;
    }

    function updateDuration(uint256 duration) public {}

    function setAdhocNBits(uint32 nbits) public onlyOwner {
        emit NBitsSet(nbits);
    }

    function setOperator(address _operator) public onlyOwner {
        operator = _operator;
        emit OperatorSet(_operator);
    }
}
