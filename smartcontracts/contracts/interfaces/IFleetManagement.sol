// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IFleetManagement {
    event MSPSet(address indexed msp);
    event CreditCenterSet(address indexed creditCenter);
    event CoordinatorSet(address indexed coordinator);
    event ProverStoreSet(address indexed prover);
    event SlasherSet(address indexed slasher);
    event FeeWithdrawn(address indexed account, uint256 amount);
    event RegistrationFeeSet(uint256 fee);

    function epoch() external view returns (uint256);
    // function msp() external view returns (address);
    // function project() external view returns (address);
    // function proverStore() external view returns (address);
    // function coordinator() external view returns (address);
    // function slasher() external view returns (address);
    function minStake() external view returns (uint256);
    function isActiveProver(uint256 _proverId) external view returns (bool);
    function isActiveCoordinator(address _coordinator, uint256 _projectId) external view returns (bool);
    function ownerOfProver(uint256 _proverId) external view returns (address);

    function grant(uint256 _proverId, uint256 _amount) external;
}
