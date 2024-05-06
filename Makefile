
DOCKER_COMPOSE_TEST_FILE=./docker-compose-test.yaml

e2e_test_depends_stop:
	@docker-compose -f ${DOCKER_COMPOSE_TEST_FILE} down

e2e_test_depends_start:
	@docker-compose -p w3bstream-sprout -f ${DOCKER_COMPOSE_TEST_FILE} up -d

.PHONY: e2e_test_depends
e2e_test_depends: e2e_test_depends_stop e2e_test_depends_start

e2e_test: e2e_test_depends
	@cd cmd/e2etest/ && go test ./... -v

unit_test:
	GOARCH=amd64 go test -gcflags="all=-N -l" `go list ./... | grep -v github.com/machinefi/sprout/cmd/e2etest` -covermode=atomic -coverprofile cover.out

.PHONY: contract_test_depends
contract_test_depends:
	@cd smartcontracts && npm install --save-dev hardhat

contract_test: contract_test_depends
	@cd smartcontracts && npx hardhat test
