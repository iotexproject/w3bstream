# w3bstream-mainnet
w3bstream mainnet

# test step
1. set your private key in docker-compose.yaml at https://github.com/machinefi/w3bstream-mainnet/blob/main/docker-compose.yaml#L18
2. run `docker-compose up`
3. run `docker logs -f w3bstream-node` to watch log
4. run `curl -X POST localhost:9000/message -H 'Content-Type: application/json' -d '{"projectID": "test01","projectVersion": "1.0", "data": "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"}'`