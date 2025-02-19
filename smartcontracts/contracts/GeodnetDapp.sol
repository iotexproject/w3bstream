// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IDapp.sol";

interface IMarshalDAOTicker {
    function tick(address _device) external;
}

contract GeodnetDapp is IDapp {
    address public verifier;
    address public ticker;

    constructor(address _verifier, address _ticker) {
        verifier = _verifier;
        ticker = _ticker;
    }

    function process(
        address _prover,
        uint256 _projectId,
        bytes32[] calldata _taskIds,
        bytes calldata _data
    ) external override {
        // Validate data length (79 uint256 values = 79 * 32 bytes)
        require(_data.length == 33 * 32, "Invalid data length");

        // Prepare function selector
        bytes4 selector = bytes4(keccak256("verifyProof(uint256[8],uint256[2],uint256[2],uint256[21])"));

        // Call verifier contract
        (bool success, ) = verifier.staticcall(abi.encodePacked(selector, _data));
        require(success, "Verifier call failed");

        bytes32[] memory data = bytesToBytes32Array(_data);
        for (uint256 i = 12; i < 32; i++) {
            bool isMoved = isBitSet(data[32], 236 + i - 12);
            if (isMoved) {
                address deviceAddr = address(bytes20(data[i]));
                IMarshalDAOTicker(ticker).tick(deviceAddr);
            }
        }
    }

    function bytesToBytes32Array(bytes memory data) public pure returns (bytes32[] memory) {
        uint256 dataNb = data.length / 32;
        bytes32[] memory dataList = new bytes32[](dataNb);
        uint256 index = 0;
        for (uint256 i = 32; i <= data.length; i = i + 32) {
            bytes32 temp;
            assembly {
                temp := mload(add(data, i))
            }
            dataList[index] = temp;
            index++;
        }
        return (dataList);
    }

    function isBitSet(bytes32 bitmap, uint256 n) public pure returns (bool) {
        require(n < 256, "n must be less than 256");
        return (uint256(bitmap) >> n) & 1 == 1;
    }
}
