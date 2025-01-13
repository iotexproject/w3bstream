// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

interface IMarshalDAOTicker {
    function tick(address _device) external;
}

contract GeodnetDapp {
    address public verifier;
    address public ticker;

    constructor(address _verifier, address _ticker) {
        verifier = _verifier;
        ticker = _ticker;
    }

    function process(
        uint256 _projectId,
        bytes32 _taskId,
        address _prover,
        address _deviceId,
        bytes calldata _data
    ) external {
        // Validate data length (79 uint256 values = 79 * 32 bytes)
        require(_data.length == 79 * 32, "Invalid data length");

        // Prepare function selector
        bytes4 selector = bytes4(keccak256("verifyProof(uint256[8],uint256[2],uint256[2],uint256[67])"));

        // Call verifier contract
        (bool success, ) = verifier.staticcall(abi.encodePacked(selector, _data));
        require(success, "Verifier call failed");

        uint256 isMoved = 0;
        for (uint256 i = 0; i < 32; i++) {
            isMoved = isMoved << 8;
            isMoved |= uint8(_data[13 * 32 + i]);
        }
        if (isMoved > 0) {
            IMarshalDAOTicker(ticker).tick(_deviceId);
        }
    }
}
