// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IFleetManagement {
    event Stake(uint256 indexed proverId, uint256 amount);
    event Unstake(uint256 indexed proverId, uint256 amount);
    event Withdrawn(uint256 indexed proverId, address indexed account, uint256 amount);
    event Grant(uint256 indexed proverId, uint256 amount);

    function epoch() external view returns (uint256);
    function project() external view returns (address);
    function prover() external view returns (address);
    function minStake() external view returns (uint256);
    function stakedAmount(uint256 _proverId) external view returns (uint256);
    function isNormalProject(uint256 _projectId) external view returns (bool);
    function isNormalProver(uint256 _proverId) external view returns (bool);

    function stake(uint256 _proverId) external payable;
    function unstake(uint256 _proverId, uint256 _amount) external;
    function withdraw(uint256 _proverId, address _to) external;
    function grant(uint256 _proverId) external payable;
}
