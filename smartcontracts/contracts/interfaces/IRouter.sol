// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IRouter {
    event DappBound(uint256 indexed projectId, address indexed operator, address dapp);
    event DappUnbound(uint256 indexed projectId, address indexed operator);

    function dapp(uint256 _projectId) external view returns (address);

    function bindDapp(uint256 _projectId, address _dapp) external;

    function unbindDapp(uint256 _projectId) external;

    function route(address _prover, uint256 _projectId, bytes32[] calldata _taskIds, bytes calldata _data) external;
}
