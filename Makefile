# Variables
BUILD_DIR=build
DOCKER_APINODE_TARGET=ghcr.io/iotexproject/w3bstream-apinode:latest
DOCKER_PROVER_TARGET=ghcr.io/iotexproject/w3bstream-prover:latest
DOCKER_BOOTNODE_TARGET=ghcr.io/iotexproject/w3bstream-bootnode:latest
DOCKER_SEQUENCER_TARGET=ghcr.io/iotexproject/w3bstream-sequencer:latest

# Build targets
.PHONY: build
build: build-apinode build-prover build-bootnode build-sequencer

.PHONY: build-apinode
build-apinode:
	go build -o $(BUILD_DIR)/apinode cmd/apinode/main.go

.PHONY: build-prover
build-prover:
	go build -o $(BUILD_DIR)/prover cmd/prover/main.go

.PHONY: build-bootnode
build-bootnode:
	go build -o $(BUILD_DIR)/bootnode cmd/bootnode/main.go

.PHONY: build-sequencer
build-sequencer:
	go build -o $(BUILD_DIR)/sequencer cmd/sequencer/main.go

# Docker targets
.PHONY: images
images:
	docker build -f build/apinode.Dockerfile -t $(DOCKER_APINODE_TARGET) .
	docker build -f build/prover.Dockerfile -t $(DOCKER_PROVER_TARGET) .
	docker build -f build/bootnode.Dockerfile -t $(DOCKER_BOOTNODE_TARGET) .
	docker build -f build/sequencer.Dockerfile -t $(DOCKER_SEQUENCER_TARGET) .

.PHONY: unit_test
unit_test:
	GOARCH=amd64 go test -gcflags="all=-N -l" ./... -covermode=atomic -coverprofile cover.out

# Clean targets
.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)

.PHONY: contract-test
contract-test:
	@cd smartcontracts && npm install --save-dev hardhat
	@cd smartcontracts && npx hardhat test

.PHONY: e2e-test
e2e-test:
	@cd smartcontracts && yarn install
	@TEST_E2E=true go test ./e2e -v
