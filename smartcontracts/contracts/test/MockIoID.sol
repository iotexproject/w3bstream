// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MockIoID is ERC721 {
    event CreateIoID(address indexed owner, uint256 id, address wallet, string did);

    uint256 _Id;
    mapping(address => uint256) public deviceProject;
    string public constant METHOD = "did:io:";

    constructor() ERC721("Mock IoID", "MIoID") {}

    function register(uint256 _projectId, address _device, string memory _deviceDID) external returns (uint256 id_) {
        id_ = ++_Id;

        _mint(msg.sender, _Id);
        deviceProject[_device] = _projectId;

        emit CreateIoID(msg.sender, id_, msg.sender, _deviceDID);
    }
}
