// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./W3bstreamProject.sol";

contract W3bstreamProjectList {

    struct Attribute {
        bytes32 k;
        bytes v;
    }

    struct Project {
        uint256 id;
        string uri;
        bytes32 hash;
        bool isPaused;
        Attribute[] attributes;
    }

    function list(address _projectContract, bytes32[] memory _attributeKeys) external view returns (uint256 blockNumber_, Project[] memory projects_) {
       W3bstreamProject w3bstreamProject = W3bstreamProject(_projectContract);
       uint256 count = w3bstreamProject.count();
       projects_ = new Project[](count);
       blockNumber_ = block.number;

        for (uint256 i = 1; i <= count; i++) {
            Project memory project;

            project.id = i;

            W3bstreamProject.ProjectConfig memory config =  w3bstreamProject.config(i);
            project.uri = config.uri;
            project.hash = config.hash;

            project.isPaused=w3bstreamProject.isPaused(i);

            Attribute[] memory attrs = new Attribute[](_attributeKeys.length);
            for (uint j = 0; j < _attributeKeys.length; j++) {
                Attribute memory attr;
                attr.k = _attributeKeys[j];
                attr.v = w3bstreamProject.attribute(i, _attributeKeys[j]);
                attrs[j] = attr;
            }
            project.attributes = attrs;

            projects_[i-1] = project;
        }
    }
}
