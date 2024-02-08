// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IFleetManager} from "./interfaces/IFleetManager.sol";

contract FleetManager is IFleetManager, Initializable {
    function initialize() public initializer {}

    function isAllowed(address _operator, uint256 _projectId) external view override returns (bool) {
        return true;
    }
}
