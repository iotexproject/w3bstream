pragma solidity ^0.5.0;

contract Store {
    bytes private proof;

    function setProof(bytes memory _proof) public {
        proof = _proof;
    }

    function getProof() public view returns (bytes memory){
        return proof;
    }
}
