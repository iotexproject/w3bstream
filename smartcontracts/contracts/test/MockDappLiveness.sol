// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MockDappLiveness {
    error CustomError();

    uint8 public errorType;
    address public verifier;

    constructor(address _verifier) {
        verifier = _verifier;
    }

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

        // Validate data length (78 uint256 values = 78 * 32 bytes)
        require(_data.length == 78 * 32, "Invalid data length");

        // Prepare function selector
        bytes4 selector = bytes4(keccak256("verifyProof(uint256[8],uint256[2],uint256[2],uint256[66])"));

        // Call verifier contract
        (bool success, ) = verifier.staticcall(abi.encodePacked(selector, _data));
        require(success, "Verifier call failed");

        // TODO: cross-validate the public inputs which are the last 66 uint256 values in the _data
    }
}
