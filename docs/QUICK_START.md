# Quick Start

## Prerequisites

Ensure you have the following tools installed:

- [ioctl](https://docs.iotex.io/builders/reference-docs/ioctl-client#install-latest-release-build): For interacting with the IoTeX blockchain.
- [Cargo](https://doc.rust-lang.org/cargo/getting-started/installation.html): Required for building the prover code.
- [curl](https://curl.se/): For sending messages to the API node.

- [jq](https://jqlang.github.io/jq/): To format JSON output.

## Step 1: Create and fund a Developer wallet

Start by creating an IoTeX developer wallet and funding it with test tokens.

```sh
ioctl account createadd devaccount
```

Note the 0x wallet address provided. You can claim test IOTX tokens on the Developer portal at <https://developers.iotex.io/faucet>

## Step 2: Register a new DePIN Project on IoTeX

To register a new project: 

```sh
ioctl ioid register "your project unique name"
```

Take note of your `Project ID` and set it as an environment variable:

```sh
export PROJECT_ID=YOUR_PROJECT_ID
```

## Step 3: Deploy the Demo prover

Next, deploy a demo range prover based on Risc0. Follow these steps:

```sh
git clone https://github.com/iotexproject/w3bstream

cd w3bstream/examples/risc0-circuit

ioctl ws project update --id $PROJECT_ID --path range_prover.json
```

## Step 4: Link Your Project to the Dummy DApp Contract

Bind your project to a sample DApp contract to ensure it receives your ZK proofs:

```sh
ioctl ws router bind --project-id $PROJECT_ID --dapp 0xba0104cD02672840Da55Bb14bebdd047Dfbfc02B
```

<details>
  <summary> Learn more</summary>
    The "dummy dapp" is a simple contract that will store ZK-Proofs from W3bstream without further processing. This can be useful for debugging proof generation if your Dapp is experiencing issues.
</details>

## Step 5: Activate Your Project

After deployment, the project is initially set to “paused,” meaning W3bstream won’t generate proofs for it until resumed:

```sh
ioctl ws project resume --id $PROJECT_ID
```

## Step 6: Test the project

The demo range prover requires a message with two inputs: an integer to test and the range to test within. Send this message to the W3bstream API node as follows:

```sh
curl --location 'https://dragonfruit-testnet.w3bstream.com/message' \
--header 'Content-Type: application/json' \
--data '{
    "projectID": '"$PROJECT_ID"',
    "projectVersion": "v1.0.0",
    "data": "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}",
    "Signature": "0x2978cfc3f4dd378c6bf4cac4b8164d1f8ce32c251b6c0c0ed5dfd48c907a82af4185e933f01e48dd8e26f84b376b7158140ebfcd2c4400c030ec2c6c0ffd37ae00"
}'
```

Record the Task ID from the response, wait a few minutes, and then check its status:

```sh
curl --location --request GET 'https://dragonfruit-testnet.w3bstream.com/task' \                                   
--header 'Content-Type: application/json' \
--data '{
    "projectID": '"$PROJECT_ID"',
    "taskID": "YOUR TASK ID HERE"
}' | jq
```

Example result:

```json
{
  "projectID": 934,
  "taskID": "0xec1d012fff25fdfe3e16e823a1d7ab1315b5fa76f19d7081724f8ee8a52c9633",
  "states": [
    {
      "state": "packed",
      "time": "2024-10-29T15:53:22.121767Z"
    },
    {
      "state": "assigned",
      "time": "2024-10-29T15:53:30.923506Z",
      "prover_id": "did:io:862FB1E46a94F62D792Ac24aeF12e02Ea6673024"
    },
    {
      "state": "settled",
      "time": "2024-10-29T15:54:10.910397Z",
      "comment": "The task has been completed. Please check the generated proof in the corresponding DApp contract.",
      "transaction_hash": "0xc27c21fb6f8a1d208680ad6c300884a14eba3525862ee8b79e1800920bc32f17"
    }
  ]
}
```

When the Task ID reaches “settled” status, you can verify the `transaction_hash` on <https://testnet.iotexscan.io> to review the proof written to the DApp contract.

## What's next??

Refer to the [Developer Guide](./DEVELOPER_GUIDE.md) to learn how to customize the prover and the Dapp contract.

## Additional commands

### Retrieve Project Info

Use the following command to retrieve information about your project:

```bash
ioctl ws project query --id "your project id"
```

### Set the Required number of Provers of the Project

By default, only one prover processes the project’s data. To customize this number, use:

```bash
ioctl ws project attributes set --id "your project id" --key "RequiredProverAmount" --val "your expected amount"
```

### Stop the Project

To stop data processing and proof generation for your project, use the command below:

```bash
ioctl ws project pause --id "your project id"
```

### Unbind the Project from the DApp

If you want to unbind the project from the DApp:

```sh
ioctl ws router unbind --project-id "your project id"
```

