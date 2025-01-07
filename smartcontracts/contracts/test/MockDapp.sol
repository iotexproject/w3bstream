// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MockDapp {
    error CustomError();

    uint8 public errorType;

    function setErrorType(uint8 _errorType) external {
        errorType = _errorType;
    }

    function process(
        uint256 _projectId,
        bytes32 _taskId,
        address _prover,
        address _deviceId,
        bytes calldata _data
    ) external view {
        if (errorType == 1) {
            require(false, "Normal Error");
        } else if (errorType == 2) {
            revert CustomError();
        }
    }
}
