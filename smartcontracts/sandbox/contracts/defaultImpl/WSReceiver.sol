// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IWSReceiver} from "../interfaces/IReceiver.sol";

contract WSReceiver is IWSReceiver {
    address public deviceNFTRegistry;

    uint256 currentBatchHeight;
    mapping(uint256 => Batch) batches;

    constructor(address _deviceNFTRegistry) {
        deviceNFTRegistry = _deviceNFTRegistry;
    }

    function receiveData(bytes calldata _data) external {
        _verify(_data[64:]);

        bytes32 _batchMR = bytes32(_data[:32]);
        bytes32 _devicesMR = bytes32(_data[32:64]);

        uint256 newHeight = currentBatchHeight + 1;
        batches[currentBatchHeight] = Batch(_batchMR, _devicesMR);
        currentBatchHeight = newHeight;
    }

    function getBatchHeight() external view returns (uint256) {
        return currentBatchHeight;
    }

    function _verify(bytes calldata) internal pure returns (bool) {
        return true;
    }
}
