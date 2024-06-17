// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;


contract VerifierRouterTemplate {

    uint256 constant bn254Prime = 21888242871839275222246405745257275088548364400416034343698204186575808495617;

    function verify(address verifier, uint256 publicInput, uint256 projectID, uint256 taskID, bytes calldata proof) view public {
        bytes32 _publicInput = uint256ToFr(publicInput);
        bytes32 _projectID = uint256ToFr(projectID);
        bytes32 _taskID = uint256ToFr(taskID);
        bytes memory callData = abi.encodePacked(_publicInput, _projectID, _taskID, proof);

        (bool success,) = verifier.staticcall(callData);
        require(success, "Failed to verify proof");
        // TODO
    }

    function uint256ToFr(uint256 _value) public pure returns (bytes32) {
        return bytes32(_value % bn254Prime);
    }
}
