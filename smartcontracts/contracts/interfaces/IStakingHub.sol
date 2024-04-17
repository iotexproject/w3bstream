// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IStakingHub {
    // event Stake(uint256 indexed proverId, uint256 amount);
    // event Unstake(uint256 indexed proverId, uint256 amount);
    // event Withdrawn(uint256 indexed proverId, address indexed account, uint256 amount);
    // event Grant(uint256 indexed proverId, uint256 amount);
    // event CoordinatorSet(address indexed coordinator);
    // event ProverStoreSet(address indexed prover);
    // event SlasherSet(address indexed slasher);

    // function stake(uint256 _proverId) external payable;
    // function unstake(uint256 _proverId, uint256 _amount) external;
    // function withdraw(uint256 _proverId, address _to) external;
    function stakedAmount(address account) external view returns (uint256);
}
