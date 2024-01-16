// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract ZNodeRegistrar  {
    mapping(string => bool) public znodes;

    event ZNodeAdded(string indexed did);
   
    function createZNode(string memory _did) public {
        znodes[_did] = true;
        
        emit ZNodeAdded(_did);
    }
}
