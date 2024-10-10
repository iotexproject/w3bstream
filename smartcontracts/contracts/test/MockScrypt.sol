// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IScrypt} from "../interfaces/IBlockHeaderValidator.sol";

contract MockScrypt {
    bytes public _hashValue;

    function setHash(bytes memory _hash) public {
        _hashValue = _hash;
    }

    function hash(
        bytes calldata,
        bytes calldata,
        uint64,
        uint32,
        uint32,
        uint32,
        uint32
    ) external view returns (bytes memory retval) {
        return _hashValue;
    }
}