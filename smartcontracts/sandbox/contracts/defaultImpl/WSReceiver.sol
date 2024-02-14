// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {IWSReceiver} from "../interfaces/IReceiver.sol";

contract WSReceiver is IWSReceiver {
    address public deviceNFTRegistry;

    mapping(uint256 => Tunnel) public tunnels;

    constructor(address _deviceNFTRegistry) {
        deviceNFTRegistry = _deviceNFTRegistry;
    }

    function receiveData(uint256 _tunnelId, bytes32 _batchMR, bytes32 _devicesMR, bytes calldata _zkProof) external {
        _verify(_zkProof);

        tunnels[_tunnelId].batches[tunnels[_tunnelId].currentBatchHeight] = Batch(_batchMR, _devicesMR);
    }

    function getBatchHeight(uint256 _tunnelId) external view returns (uint256) {
        return tunnels[_tunnelId].currentBatchHeight;
    }

    function _verify(bytes calldata) internal pure returns (bool) {
        return true;
    }
}
