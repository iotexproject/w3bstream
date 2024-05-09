
DOCKER_COMPOSE_TEST_FILE=./docker-compose-test.yaml

e2e_test_depends_stop:
	@docker-compose -f ${DOCKER_COMPOSE_TEST_FILE} down

e2e_test_depends_start:
	@docker-compose -p w3bstream-sprout -f ${DOCKER_COMPOSE_TEST_FILE} up -d

.PHONY: e2e_test_depends
e2e_test_depends: e2e_test_depends_stop e2e_test_depends_start

e2e_test: e2e_test_depends
	@cd cmd/e2etest/ && CGO_LDFLAGS='-L../sequencer/lib/linux-x86_64 -lioConnectCore' go test ./... -v

unit_test:
	GOARCH=amd64 go test -gcflags="all=-N -l" `go list ./... | grep -v github.com/machinefi/sprout/cmd/e2etest` -covermode=atomic -coverprofile cover.out

.PHONY: contract_test_depends
contract_test_depends:
	@cd smartcontracts && npm install --save-dev hardhat

contract_test: contract_test_depends
	@cd smartcontracts && npx hardhat test

.PHONY: images
images:
	@for target in 'sequencer' 'prover' 'coordinator' ;                \
	do                                                                 \
		echo build $$target image ;                                    \
		if [ -e $$target.Dockerfile ]; then                            \
			echo $$target.Dockerfile ;                                 \
			docker build -f $$target.Dockerfile . -t $(USER)/$$target; \
		else                                                           \
			echo "no entry";                                           \
		fi;                                                            \
		echo "done!";                                                  \
	done


MOD=$(shell cat go.mod | grep ^module -m 1 | awk '{ print $$2; }' || '')

.PHONY: fmt
fmt:
	@for item in `find . -type f -name '*.go' -not -path '*.pb.go'` ; \
    do \
		if [ -z $$MOD ]; then \
			goimports -w $$item ; \
		else \
			goimports -w -local "${MOD}" $$item ; \
		fi \
    done

