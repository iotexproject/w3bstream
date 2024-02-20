// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";

import {INodeRegistry} from "./interfaces/INodeRegistry.sol";

contract NodeRegistry is INodeRegistry, ERC721Upgradeable {
    uint256 internal nextTokenId;
    mapping(uint256 => Node) internal _nodes;
    mapping(address => uint256) internal _operators;

    function initialize() public initializer {
        __ERC721_init("W3bstream node registry", "WNR");
        nextTokenId = 0;
    }

    function register(address _operator) external override {
        if (_operator == address(0)) {
            revert InvalidAddress();
        }

        ++nextTokenId;
        _nodes[nextTokenId] = Node(nextTokenId, true, _operator);
        _operators[_operator] = nextTokenId;
        _safeMint(msg.sender, nextTokenId);

        emit NodeRegistered(msg.sender, nextTokenId, _operator);
    }

    function updateOperator(uint256 _tokenId, address _operator) external override {
        if (ownerOf(_tokenId) != msg.sender) {
            revert NotNodeOwner();
        }
        if (_operator == address(0)) {
            revert InvalidAddress();
        }
        if (_operators[_operator] != 0) {
            revert OperatorAlreadyRegistered();
        }

        _nodes[_tokenId].operator = _operator;
        _operators[_operator] = _tokenId;
        emit NodeUpdated(_tokenId, _operator);
    }

    function getNode(uint256 _tokenId) external view override returns (Node memory) {
        return _nodes[_tokenId];
    }

    function getNodeByOperator(address _operator) external view override returns (Node memory) {
        uint256 _tokenId = _operators[_operator];
        if (_tokenId == 0) {
            revert OperatorUnregister();
        }
        return _nodes[_tokenId];
    }
}
