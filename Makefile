
DOCKER_COMPOSE_TEST_FILE=./docker-compose-test.yaml

integration_test_depends_stop:
	@docker-compose -f ${DOCKER_COMPOSE_TEST_FILE} down

integration_test_depends_start:
	@docker-compose -p w3bstream-sprout -f ${DOCKER_COMPOSE_TEST_FILE} up -d

.PHONY: integration_test_depends
integration_test_depends: integration_test_depends_stop integration_test_depends_start

integration_test: integration_test_depends
	@cd cmd/tests/ && go test ./... -v

unit_test:
	go test -p 1 -gcflags="all=-N -l" `go list ./... | grep -v github.com/machinefi/sprout/cmd/tests` -covermode=atomic -coverprofile cover.out
