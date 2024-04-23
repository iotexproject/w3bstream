// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract MultiCall {
    function multiCall(address[] calldata _targets, bytes[] calldata _data) external view returns (bytes[] memory results_) {
        require(_targets.length == _data.length, "target length != data length");

        results_ = new bytes[](_data.length);

        for (uint256 i = 0; i < _targets.length; i++) {
            (bool success, bytes memory result) = _targets[i].staticcall(_data[i]);
            if (!success) {
                break;
            }
            results_[i] = result;
        }
    }
}
