// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IRouter {
    event DataProcessed(address indexed operator, bool success, string revertReason);
    event BindDapp(uint256 indexed projectId, address indexed operator, address dapp);
    event UnbindDapp(uint256 indexed projectId, address indexed operator);

    function fleetManagement() external view returns (address);
    function dapp(uint256 _projectId) external view returns (address);

    function bindDapp(uint256 _projectId, address _dapp) external;
    function unbindDapp(uint256 _projectId) external;
    function route(uint256 _projectId, uint256 _proverId, bytes calldata _data) external;
}
