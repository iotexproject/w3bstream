

integration_test_depends_stop:
	@docker stop postgres_test || true && docker container rm postgres_test || true
	@docker stop didkit_test || true && docker container rm didkit_test || true
	@docker stop halo2_test || true && docker container rm halo2_test || true
	@docker stop risc0_test || true && docker container rm risc0_test || true
	@docker stop zkwasm_test || true && docker container rm zkwasm_test || true
	@docker stop wasm_test || true && docker container rm wasm_test || true

.PHONY: integration_test_depends
integration_test_depends: integration_test_depends_stop postgres_test didkit_test halo2_test risc0_test zkwasm_test wasm_test

.PHONY: postgres_test
postgres_test:
	docker run --name postgres_test \
  -e POSTGRES_USER=test_user \
  -e POSTGRES_PASSWORD=test_passwd \
  -e POSTGRES_DB=test \
  -p 15432:5432 \
  -d postgres:14

.PHONY: didkit_test
didkit_test:
	docker run --name didkit_test \
  -e DIDKIT_HTTP_HTTP_PORT=19999 \
  -e DIDKIT_HTTP_HTTP_ADDRESS='[0,0,0,0]' \
  -e DIDKIT_HTTP_HTTP_BODYSIZELIMIT='2097152' \
  -e DIDKIT_HTTP_ISSUER_KEYS='[{"kty":"OKP","crv":"Ed25519","x":"THRnf4Zj7gh93XTnII8G0tQIoYb4IbkoTqcy5TNKJlg","d":"es8N8nmdU9o5wWdCEMc2xKCigN1LKc6xro1efDy7Y7M"}, {"kty":"OKP","crv":"Ed25519","x":"STSryIxBN3pyC5YQ5GnjlMmILUWcb5M0_sHpqxxmsog","d":"BiKwVOhhI-fcMMjfcxo2AdB3ygamMmgcMzjaOUl7O6s"}]' \
  -p 19999:9999 \
  -d ghcr.io/spruceid/didkit-http:latest

.PHONY: risc0_test
risc0_test:
	docker run --name risc0_test \
  --platform linux/x86_64 \
  -e DATABASE_URL=postgres://test_user:test_passwd@localhost:15432/test?sslmode=disable \
  -e BONSAI_URL: https://api.bonsai.xyz} \
  -e BONSAI_KEY= xxx \
  -p 14001:4001 \
  -d wangweixiaohao2944/risc0server:v1.0.0.rc4

.PHONY: halo2_test
halo2_test:
	docker run --name halo2_test \
  --platform linux/x86_64 \
  -p 14002:4002 \
  -d wangweixiaohao2944/halo2server:v0.0.6

.PHONY: zkwasm_test
zkwasm_test:
	docker run --name zkwasm_test \
	--platform linux/x86_64 \
	-p 14003:4003 \
	-d iotexdev/zkwasmserver:v0.0.3

.PHONY: wasm_test
wasm_test:
	docker run --name wasm_test \
	--platform linux/x86_64 \
	-p 14004:4004 \
	-d wangweixiaohao2944/wasmserver:v0.0.1.rc0

integration_test: integration_test_depends
	@cd cmd/test/ && go test ./... -v

unit_test:
	go test -p 1 -gcflags="all=-N -l" `go list ./... | grep -v github.com/machinefi/sprout/cmd/tests` -covermode=atomic -coverprofile cover.out
