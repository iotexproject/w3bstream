## Interacting with W3bstream node

Install **ioctl**: The command-line interface for interacting with the IoTeX blockchain.

```bash
brew tap iotexproject/ioctl-unstable
brew install iotexproject/ioctl-unstable/ioctl-unstable
alias ioctl=`which ioctl-unstable`
```

[More on the IoTeX ioctl client →](https://docs.iotex.io/the-iotex-stack/wallets/command-line-client)

### Sending messages

Send a message to a RISC0-based test project (ID 10000):

```bash
ioctl ws message send --project-id 10000 --project-version "0.1" --data "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"
```

Send a message to the Halo2-based test project (ID 10001):

```bash
ioctl ws message send --project-id 10001 --project-version "0.1" --data "{\"private_a\": 3, \"private_b\": 4}"
```

Send a message to a zkWasm-based test project (ID 10002):

```bash
ioctl ws message send --project-id 10002 --project-version "0.1" --data "{\"private_input\": [1, 1] , \"public_input\": [] }"
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
