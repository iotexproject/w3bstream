// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MockDapp {
    error CustomError();

    uint8 public errorType;

    function setErrorType(uint8 _errorType) external {
        errorType = _errorType;
    }

    function process(
        address _prover,
        uint256 _projectId,
        bytes32[] calldata _taskIds,
        bytes calldata _data
    ) external view {
        if (errorType == 1) {
            require(false, "Normal Error");
        } else if (errorType == 2) {
            revert CustomError();
        }
    }
}
