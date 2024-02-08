// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/// @title Sandbox data receiver
interface IReceiver {
    /// @notice device nft registry
    /// @return address of device NFT registry
    function deviceNFTRegistry() external view returns (address);

    /// @notice receive prover data
    /// @param _data .
    function receiveData(bytes memory _data) external;
}
