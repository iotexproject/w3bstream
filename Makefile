
DOCKER_COMPOSE_TEST_FILE=./docker-compose-test.yaml

integration_test_depends_stop:
	@docker-compose -f ${DOCKER_COMPOSE_TEST_FILE} down

integration_test_depends_start:
	@docker-compose -p w3bstream-sprout -f ${DOCKER_COMPOSE_TEST_FILE} up -d

.PHONY: integration_test_depends
integration_test_depends: integration_test_depends_stop integration_test_depends_start

integration_test: integration_test_depends
	@cd cmd/test/ && go test ./... -v

unit_test:
	GOARCH=amd64 go test -gcflags="all=-N -l" `go list ./... | grep -v github.com/machinefi/sprout/cmd/test` -covermode=atomic -coverprofile cover.out

.PHONY: contract_test_depends
contract_test_depends:
	@cd smartcontracts && npm install --save-dev hardhat

contract_test: contract_test_depends
	@cd smartcontracts && npx hardhat test
