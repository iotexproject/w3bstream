// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MockIoID is ERC721 {
    uint256 _Id;

    constructor() ERC721("Mock IoID", "MIoID") {}

    function register() external returns (uint256) {
        ++_Id;

        _mint(msg.sender, _Id);
        return _Id;
    }
}
