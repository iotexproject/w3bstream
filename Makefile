.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: docker
docker:
	DOCKER_BUILDKIT=1 docker build -t $(USER)/w3bstream-mainnet-node:local .