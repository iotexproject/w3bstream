// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {IFleetManager} from "./interfaces/IFleetManager.sol";
import {INodeRegistry} from "./interfaces/INodeRegistry.sol";

contract FleetManager is IFleetManager, Initializable {
    address public projectRegistry;
    address public nodeRegistry;

    mapping(uint256 => mapping(address => bool)) internal _nodes;

    function initialize(address _projectRegistry, address _nodeRegistry) public initializer {
        projectRegistry = _projectRegistry;
        nodeRegistry = _nodeRegistry;
    }

    modifier onlyProjectOwner(uint256 _projectId) {
        if (IERC721(projectRegistry).ownerOf(_projectId) != msg.sender) {
            revert NotProjectOwner();
        }
        _;
    }

    function allow(uint256 _projectId, address _node) external override onlyProjectOwner(_projectId) {
        if (_nodes[_projectId][_node]) {
            revert NodeAlreadyAllowed();
        }

        _nodes[_projectId][_node] = true;
        emit NodeAllowed(_projectId, _node);
    }

    function disallow(uint256 _projectId, address _node) external override onlyProjectOwner(_projectId) {
        if (!_nodes[_projectId][_node]) {
            revert NodeNotAllow();
        }

        emit NodeDisallowed(_projectId, _node);
    }

    function isAllowed(address _operator, uint256 _projectId) external view returns (bool) {
        if (_operator == address(0)) {
            revert InvalidOperatorAddress();
        }

        address _node = _getOperatorNode(_operator);
        if (_node == address(0)) {
            revert NodeUnregister();
        }

        return _nodes[_projectId][_node];
    }

    function _getOperatorNode(address _operator) internal view returns (address) {
        return INodeRegistry(nodeRegistry).getNode(_operator).node;
    }
}
