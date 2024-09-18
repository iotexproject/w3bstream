// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MockProcessor {
    error CustomError();

    uint8 public errorType;

    function setErrorType(uint8 _errorType) external {
        errorType = _errorType;
    }

    function process(uint256, uint256, string memory, bytes memory) external view {
        if (errorType == 1) {
            require(false, "Normal Error");
        } else if (errorType == 2) {
            revert CustomError();
        }
    }
}
