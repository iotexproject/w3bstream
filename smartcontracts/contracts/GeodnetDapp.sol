// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./interfaces/IDapp.sol";

interface IMarshalDAOTicker {
    function tick(address _device) external;
}

contract GeodnetDapp is IDapp {
    address public ticker;

    constructor(address _ticker) {
        ticker = _ticker;
    }

    function process(
        address _prover,
        uint256 _projectId,
        bytes32[] calldata _taskIds,
        bytes calldata _data
    ) external override {
        require(_data.length == 32, "invalid _data length");
        address deviceAddr = abi.decode(_data, (address));
        IMarshalDAOTicker(ticker).tick(deviceAddr);
    }
}
