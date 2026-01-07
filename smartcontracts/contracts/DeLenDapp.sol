// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;
import "./interfaces/IDapp.sol";

interface IMarshalDAOTicker {
    function tick(address _device) external;
}

contract DelenDapp is IDapp {
    address public verifier;
    mapping(bytes32 => uint64) public counter;

    constructor(address _verifier) {
        verifier = _verifier;
    }

    function process(
        address _prover,
        uint256 _projectId,
        bytes32[] calldata _taskIds,
        bytes calldata _data
    ) external override {
        require(_data.length == 14 * 32, "Invalid data length");

        // Prepare function selector
        bytes4 selector = bytes4(keccak256("verifyProof(uint256[8],uint256[2],uint256[2],uint256[2])"));

        // Call verifier contract
        (bool success, ) = verifier.staticcall(abi.encodePacked(selector, _data));
        require(success, "Verifier call failed");

        bytes32[] memory data = bytesToBytes32Array(_data);
        // TODO: extract ioID from the proof
        bytes32 ioID = data[13];
        counter[ioID]++;
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
}
