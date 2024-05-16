// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MockProject is ERC721 {
    uint256 _projectId;

    constructor() ERC721("Mock Project", "MPN") {}

    function register() external returns (uint256) {
        ++_projectId;

        _mint(msg.sender, _projectId);
        return _projectId;
    }
}
