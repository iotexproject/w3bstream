## Interacting with W3bstream node
Install **jq**:
```bash
brew install jq
```

Install **ioctl**: The command-line interface for interacting with the IoTeX blockchain.

```bash
brew tap iotexproject/ioctl-unstable
brew install iotexproject/ioctl-unstable/ioctl-unstable
alias ioctl=`which ioctl-unstable`
```

[More on the IoTeX ioctl client →](https://docs.iotex.io/the-iotex-stack/wallets/command-line-client)

### Sending messages
Examples of sending messages to pre-created projects:

Send a message to a RISC0-based test project (ID 1):

```bash
ioctl ws message send --project-id 1 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```

Send a message to the Halo2-based test project (ID 2):

```bash
ioctl ws message send --project-id 2 --project-version "0.1" --data "{\"private_a\": 3, \"private_b\": 4}"
```

Send a message to a zkWasm-based test project (ID 3):

```bash
ioctl ws message send --project-id 3 --project-version "0.1" --data "{\"private_input\": [1, 1] , \"public_input\": [] }"
```

### Query the status of a proof request

After sending a message, you'll receive a message ID as a response from the node, e.g.,

```json
{
 "messageID": "4abbc43a-798f-49e8-bc05-b6baeafec630"
}
```

you can query the status of the message request with:

```bash
ioctl ws message query --message-id "4abbc43a-798f-49e8-bc05-b6baeafec630"
```

example result:

```json
{
 "messageID": "4abbc43a-798f-49e8-bc05-b6baeafec630",
 "states": [{
   "state": "received",
   "time": "2023-12-06T16:11:03.498785+08:00",
   "comment": ""
  },
  {
   "state": "fetched",
   "time": "2023-12-06T16:11:04.663608+08:00",
   "comment": ""
  },
  {
   "state": "proving",
   "time": "2023-12-06T16:11:04.664008+08:00",
   "comment": ""
  }
 ]
}
```

When the request is in "proved" state, you can check out the comment to find out the hash of the blockchain transaction that wrote the proof to the destination chain.

### Get verifiable credential token (WIP)
> NOTE: The following mock client DID, which have already been binded to existing projects, as $CLIENT_ID to get a VC token.
>> did:example:d23dd687a7dc6787646f2eb98d0 
>> did:key:z6MkeeChrUs1EoKkNNzoy9FwJJb9gNQ92UT8kcXZHMbwj67B
>> did:ethr:0x9d9250fb4e08ba7a858fe7196a6ba946c6083ff0 

Assuming that we are going to interact with W3bstream Sprout Staging server (`SERVER=http://sprout-staging.w3bstream.com`), and env `CLIENT_DID` has been set, the following command is used to exchange a DID token with server:

```bash
export DID_TOKEN=`echo '{
  "credential": {
    "@context": "https://www.w3.org/2018/credentials/v1",
    "id": "http://example.org/credentials/3731",
    "type": [
      "VerifiableCredential"
    ],
    "issuer": "did:key:z6MkjP2Pa1pkUgz2rP6yTXpATe4qd7ahwsGAQuU697JpcCLf",
    "issuanceDate": "2020-08-19T21:41:50Z",
    "credentialSubject": {
      "id": "'$CLIENT_DID'"
    }
  },
  "options": {
    "verificationMethod": "did:key:z6MkjP2Pa1pkUgz2rP6yTXpATe4qd7ahwsGAQuU697JpcCLf#z6MkjP2Pa1pkUgz2rP6yTXpATe4qd7ahwsGAQuU697JpcCLf",
    "proofPurpose": "assertionMethod",
    "proofFormat": "jwt"
  }
}' | http post $SERVER:9000/sign_credential | jq -r '.verifiableCredential'`
```
