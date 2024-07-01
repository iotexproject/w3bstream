# Interacting with W3bstream

---

## Prerequisites

- jq:
  [installation instructions →](https://jqlang.github.io/jq/)

- curl:
  [installation instructions →](https://curl.se/)

## Send messages use curl

We have created preset test projects for three types of zero-knowledge proof vm
(`risc0`, `halo2` and `zkWASM`). You can interact with the **sprout**
service by _task submitting_ and _task querying_ APIs to commit proof task and
retrieve task proof status.

Examples of sending messages to pre-created projects:

Send a message to a RISC0-based test project (ID 91):

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"projectID": 91,"projectVersion": "0.1","data": "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"}'\
  https://sprout-testnet.w3bstream.com/message
```

Send a message to the Halo2-based test project (ID 92):

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"projectID": 92,"projectVersion": "0.1","data": "{\"private_a\": 3, \"private_b\": 4}"}' \
  https://sprout-testnet.w3bstream.com/message
```

Send a message to a zkWasm-based test project (ID 93):

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"projectID": 93,"projectVersion": "0.1","data": "{\"private_input\": [1, 1] , \"public_input\": [2] }"}' \
  https://sprout-testnet.w3bstream.com/message
```

## Query the status of a proof request

After sending a message, you'll receive a message ID(an uuid to identify the
unique task) as a response from the **sprout** service.

```json
{
  "messageID": "8785a42c-9d6c-4780-910c-de0147aea243"
}
```

you can query the status of the message request with:

```bash
curl https://sprout-testnet.w3bstream.com/message/8785a42c-9d6c-4780-910c-de0147aea243 | jq -r '.'
```

example result:

```json
{
  "messageID": "8785a42c-9d6c-4780-910c-de0147aea243",
  "states": [
    {
      "state": "received",
      "time": "2024-06-10T09:30:05.790151Z",
      "comment": "",
      "result": ""
    },
    {
      "state": "packed",
      "time": "2024-06-10T09:30:05.793218Z",
      "comment": "",
      "result": ""
    },
    {
      "state": "dispatched",
      "time": "2024-06-10T09:30:10.87987Z",
      "comment": "",
      "result": ""
    },
    {
      "state": "proved",
      "time": "2024-06-10T09:30:11.193027Z",
      "comment": "",
      "result": "proof result"
    },
    {
      "state": "outputted",
      "time": "2024-06-10T09:30:11.20942Z",
      "comment": "output type: stdout",
      "result": ""
    }
  ]
}
```

When the request is in "proved" state, you can check out the comment to find out
the hash of the blockchain transaction that wrote the proof to the destination
chain.

## Send messages with token (Experimental)

For security purposes, **sprout** integrates ioID identity verification and
DID-Comm message encryption features. Next, you can use the `didctl`
command-line tool to simulate message encryption, allowing you to submit
encrypted proof tasks.

> note: currently, `didctl` is an experimental tool and is only supported on
> Linux. Adaptation for other os is still under development.

### install `didctl` command-line

```bash
git clone git@github.com:machinefi/ioconnect-go.git
make targets && mv cmd/didctl __YOUR_SYSTEM_PATH__
```

[More on the IoTeX ioctl client →](https://docs.iotex.io/the-iotex-stack/wallets/command-line-client)

### fetch **sprout** service did document

```bash
curl https://sprout-testnet.w3bstream.com/didDoc
```

For convenience, you can set the did document of sprout service as an
environment variable.

```bash
export serverdoc='{"@context":["https://www.w3.org/ns/did/v1","https://w3id.org/security#keyAgreementMethod"],"id":"did:io:0x81a3864898d6098b15eff17b6452fc4e28e05983","authentication":["did:io:0x81a3864898d6098b15eff17b6452fc4e28e05983#Key-secp256k1-2147483616"],"keyAgreement":["did:io:0xaefe2f283b262978a1cabc483410593d62c9c732#Key-p256-2147483617"],"verificationMethod":[{"id":"did:io:0xaefe2f283b262978a1cabc483410593d62c9c732#Key-p256-2147483617","type":"JsonWebKey2020","controller":"did:io:0x81a3864898d6098b15eff17b6452fc4e28e05983","publicKeyJwk":{"crv":"P-256","x":"xaKC13yoR2Q6FSF6mrm027-onSs9qud4OApuIE6eFd4","y":"PQk3EoMlKYf9FqorTUN8slXpNSpHyhZdxDBJ9dJmnzE","d":"","kty":"EC","kid":"Key-p256-2147483617"}},{"id":"did:io:0x81a3864898d6098b15eff17b6452fc4e28e05983#Key-secp256k1-2147483616","type":"JsonWebKey2020","controller":"did:io:0x81a3864898d6098b15eff17b6452fc4e28e05983","publicKeyJwk":{"crv":"secp256k1","x":"CBlqq_7ZfcFALq4UL-GRMrKok8Zj8XNRBCWG4XT4sVQ","y":"SopcvJFTWw8hOEUl_eE96YIcpDttqeRZSMkz4-dho6Q","d":"","kty":"EC","kid":"Key-secp256k1-2147483616"}}]}'
export serverdid=did:io:0x81a3864898d6098b15eff17b6452fc4e28e05983
```

### set simulate device client did and JWK secret envs

In the next steps, we will use a simulated device that already has an **ioID**
identity to submit encrypted request message to **sprout**. Of course, if you
already have an ioID identity and the corresponding JWK keys, you can replace
the client information below.

```bash
export clientdid=did:io:0xba80b710f0c27c8b3b72df63861e2ecea9c5aa73
export clientsec=vebfEf+v2rLUzFm2mMH9XzPbZJFzaEj3nctUCnoAbMw=
```

### request token from **sprout** service

First, use the `curl` command to request the issuance of a token from the
**sprout** service.

```bash
curl -X POST -d '{"clientID":"'$clientdid'"}' https://sprout-testnet.w3bstream.com/issue_vc
```

For security purposes, **sprout** will respond with an encrypted token. For
convenience, you can set the responded encrypted token as an env var

```bash
export ciphertoken= # response above
```

Next, we need to use the `didctl` tool to decrypt the message and obtain the
plaintext token.

```bash
didctl decrypt --cipher $ciphertoken --encryptor $serverdid --recipient $clientsec --recipient-id 2
```

For convenience, you can set the plain token as env var.

```bash
export TOKEN= #output above
```

### commit encrypted proof message

First, prepare plain request body

> note: the simulated device is bound with project 21, if you use your
> own `ioID` and JWK key to replace the `projectID` and other request
> information.

```bash
export plain_task='{"projectID": 21, "projectVersion": "0.1", "data": "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"}'
```

Next, encrypt request body

```bash
didctl encrypt --recipient $serverdoc --encryptor $clientdid --plain $plain_task
```

For convenience, you can set the encrypted request body as env var.

```bash
export cipherdata= #output above
```

Commit proof message and set

```bash
curl -X POST -d $cipherdata \
  -H "Authorization: Bearer $TOKEN" https://sprout-testnet.w3bstream.com/message
export cipherresponse= # response above
```

Last, decrypt **sprout** response to retrieve message id

```bash
didctl decrypt --recipient $clientsec --encryptor $serverdid --cipher $cipherresp --recipient-id 2
```
