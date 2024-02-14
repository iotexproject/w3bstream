// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @title IWSReceiver
interface IWSReceiver {
    struct Tunnel {
        uint256 currentBatchHeight;
        mapping(uint256 => Batch) batches;
    }

    struct Batch {
        bytes32 merkleRoot;
        bytes32 devicesMerkleRoot;
    }

    /// @notice device nft registry
    /// @return address of device NFT registry
    function deviceNFTRegistry() external view returns (address);

    /// @notice receive data to process from WSRouter
    /// @param _tunnelId tunnel id
    /// @param _batchMR merkle root of the batch
    /// @param _devicesMR merkle root of the devices
    /// @param _zkProof zk proof
    function receiveData(uint256 _tunnelId, bytes32 _batchMR, bytes32 _devicesMR, bytes calldata _zkProof) external;

    /// @notice get batch height
    function getBatchHeight(uint256 _tunnelId) external view returns (uint256);
}
