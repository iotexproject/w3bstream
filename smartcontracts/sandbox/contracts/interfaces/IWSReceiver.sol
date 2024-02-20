// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @title IWSReceiver
interface IWSReceiver {
    struct Batch {
        bytes32 merkleRoot;
        bytes32 devicesMerkleRoot;
    }

    /// @notice device nft registry
    /// @return address of device NFT registry
    function deviceNFTRegistry() external view returns (address);

    /// @notice receive data to process from WSRouter
    /// @param _data:
    /// _batchMR merkle root of the batch
    /// _devicesMR merkle root of the devices
    /// _zkProof zk proof
    function receiveData(bytes calldata _data) external;

    /// @notice get batch height
    function getBatchHeight() external view returns (uint256);
}
