.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: test
test:
	curl -X POST localhost:9000/message -H 'Content-Type: application/json' -d '{"projectID": "test01","projectVersion": "1.0", "data": "{\"private_input\":\"14\", \"public_input\":\"3,34\", \"receipt_type\":\"Snark\"}"}'

.PHONY: docker
docker:
	DOCKER_BUILDKIT=1 docker build -t $(USER)/w3bstream-mainnet-node:local .