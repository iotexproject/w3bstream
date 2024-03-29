// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IProjectRegistrar {
    event ProjectRegister(uint256 indexed projectId);
    event RegistrationFeeSet(uint256 fee);
    event WithdrawnFee(address indexed account, uint256 amount);

    function registrationFee() external view returns (uint256);

    function register(string calldata _uri, bytes32 _hash) external payable returns (uint256 _projectId);
}
