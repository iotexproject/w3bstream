// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ITaskManager, TaskAssignment} from "../interfaces/ITaskManager.sol";

contract MockTaskManager is ITaskManager {
    function assign(TaskAssignment[] calldata assignments, address, uint256) external override {}
}
