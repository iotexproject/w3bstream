// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract ProjectRegistry is ERC721 {
    uint256 internal nextTokenId;

    constructor() ERC721("Test Project Registry", "") {
        nextTokenId = 0;
    }

    function createProject(string memory, bytes32) public {
        ++nextTokenId;
        _safeMint(msg.sender, nextTokenId);
    }
}
