# Contribution Guidelines for W3bstream

## Introduction

Thank you for considering contributing to the IoTeX W3bstream! We value your contributions and want to make sure that your efforts align with the goals and standards of our project. This document provides guidelines to ensure a smooth and effective collaboration process.

## How to Contribute

There are many ways to contribute to this project:

1. **Code Contributions**: If you're looking to add or fix something in the codebase, please follow the steps outlined in the Getting started below.
2. **Bug Reports**: If you find a bug, please open an issue using the **Bug Report** template..
3. **Feature Suggestions**: Have ideas for new features? Open an issue using the **Feature Request** template.
4. **Documentation**: Improvements or additions to our documentation are always welcome. We currently use GitBook for our documentation, and you can read it at [docs.iotex.io](https://docs.iotex.io). Just locate the "Edit this Page on GitHub" link on any page to start contributing.

## Getting Started

1. **Fork this Repository**: Start by [forking the project repository](https://github.com/machinefi/sprout/fork) to your GitHub account.

2. **Clone the Repository**: Clone your forked repository to your local machine.

    ```bash
    git clone https://github.com/your-username/sprout.git
    ```

3. **Create a New Branch**: Create a branch for your changes.

    ```bash
    git checkout -b feature/your-new-feature
    ```

### Making changes

1. Ensure all required env variables are exported:
    ```bash
    # coordinator env

    export HTTP_SERVICE_ENDPOINT=:9000
    export DATABASE_DSN=postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable
    export BOOTNODE_MULTIADDR="/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o"
    export IOTEX_CHAINID=2
    export PROJECT_FILE_DIRECTORY=./test/data
    export CHAIN_ENDPOINT=https://babel-api.testnet.iotex.io
    export PROJECT_CONTRACT_ADDRESS=0x184C72E39a642058CCBc369485c7fd614B40a03d
    ```

    ```bash
    # prover env
    # --- Edit the following

    # The RPC of the destination chain where proofs must be sent
    export CHAIN_ENDPOINT=https://babel-api.testnet.iotex.io
    # The contract address to which the proof will be sent
    export PROJECT_CONTRACT_ADDRESS=0x184C72E39a642058CCBc369485c7fd614B40a03d
    # A funded account on the destination chain 
    export OPERATOR_PRIVATE_KEY=<your_blockchain_key>
    # Optional: Required for working with RISC0 provers
    export BONSAI_KEY=<your_bonsai_api_key>
    # ---
    
    export PROJECT_FILE_DIRECTORY=./test/data
    export DATABASE_URL=postgres://test_user:test_passwd@0.0.0.0:5432/test?sslmode=disable
    export ZKWASM_SERVER_ENDPOINT=localhost:4003
    export HALO2_SERVER_ENDPOINT=localhost:4002
    export RISC0_SERVER_ENDPOINT=localhost:4001
    export BOOTNODE_MULTIADDR="/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o"
    export IOTEX_CHAINID=2
    export CHAIN_CONFIG='[{"chainID":4690,"name":"iotex-testnet","endpoint":"https://babel-api.testnet.iotex.io"},{"name":"solana-testnet","endpoint":"https://api.testnet.solana.com"}]'
    ```

3. Start required services:
    
    ```bash
    docker compose -f docker-compose-dev.yaml up -d
    ```

4. Start coordinator server:
    
    ```bash
    cd cmd/coordinator && go run .
    ```

5. Open the repository in VS Code and create a launch.json file with the following content:

    ```json
    {
      "version": "0.2.0",
      "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "cwd": "${workspaceFolder}",
            "program": "${workspaceFolder}/cmd/prover"
        }
      ]
    }
    ```

6. Set your breakpoints in the code and start debugging by pressing F5! FInd the node log in the Debug Console of VS Code. 

### Guidelines

1. **Code Guidelines**: Write clean, maintainable code. Refer to the [go coding styleguide](https://google.github.io/styleguide/go/) and follow coding standards already in place.

2. **Testing**: Add tests for new features or fix existing test cases as necessary.

3. **Documentation**: Update the README, or other documentation as necessary.

### Submitting a Pull Request

1. **Create a Pull Request (PR)**: Go to this [project repository](https://github.com/machinefi/sprout) and click on 'New Pull Request'. Compare branches and create the PR from your feature branch to the main project's branch.

2. **Describe Your Changes**: In the PR description, clearly outline what you've done. Link any relevant issues.

## PR Review Process

Once you submit a PR, the project maintainers will review your changes. They may request further changes or give feedback. Keep an eye on your PR for any comments.

## Code of Conduct

We expect all contributors to be Respectful as well as considerate behavior is expected from all community members.

## Questions or Need Help?

If you have questions or need help with making contributions, feel free to reach out to us via email at <developers@iotex.io> or reach out to us on our [Discord server](https://iotex.io/devdiscord).
