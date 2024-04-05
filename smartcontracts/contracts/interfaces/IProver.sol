// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IProver {
    event OperatorSet(uint256 indexed id, address indexed operator);
    event NodeTypeUpdated(uint256 indexed id, uint256 typ);
    event ProverPaused(uint256 indexed id);
    event ProverResumed(uint256 indexed id);

    function count() external view returns (uint256);
    function nodeType(uint256 _id) external view returns (uint256);
    function operator(uint256 _id) external view returns (address);
    function isPaused(uint256 _id) external view returns (bool);
    function ownerOfOperator(address _operator) external view returns (uint256, address);

    function register(uint256 _type, address _operator) external returns (uint256);
    function changeOperator(uint256 _id, address _operator) external;
    function updateNodeType(uint256 _id, uint256 _type) external;
    function pause(uint256 _id) external;
    function resume(uint256 _id) external;
}
