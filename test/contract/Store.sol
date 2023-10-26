pragma solidity ^0.5.0;

contract Store {
    string private proof;

    function setProof(string memory _proof) public {
        proof = _proof;
    }

    function getProof() public view returns (string memory){
        return proof;
    }
}
