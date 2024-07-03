// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IRouter {
    event DataProcessed(
        uint256 indexed projectId,
        uint256 indexed router,
        address indexed operator,
        bool success,
        string revertReason
    );
    event DappBound(uint256 indexed projectId, address indexed operator, address dapp);
    event DappUnbound(uint256 indexed projectId, address indexed operator);

    // function fleetManagement() external view returns (address);
    function dapp(uint256 _projectId) external view returns (address);
    // function credits(uint256 _proverId) external view returns (uint256);

    function bindDapp(uint256 _projectId, address _dapp) external;
    function unbindDapp(uint256 _projectId) external;
    function route(uint256 _projectId, uint256 _proverId, uint256 _taskId, bytes calldata _data) external;
}
