.DEFAULT_GOAL := update

.PHONY: update
update: fetch generate_abi generate_go

.PHONY: fetch
fetch:
	@echo "#### clone ioID-contracts main branch..."
	@if [ -d "ioid" ]; then \
		cd ioid && git pull --quiet origin main; \
	else \
		git clone -b main --quiet git@github.com:machinefi/ioID-contracts.git ioid; \
	fi
	@echo DONE

.PHONY: generate_abi
generate_abi: generate_abi_ioid generate_abi_project_device

.PHONY: generate_abi_ioid
generate_abi_ioid:
	@echo "#### generate abis of ioID from latest contracts..."
	@cd ioid && yarn install > /dev/null 2>&1 && yarn hardhat compile > /dev/null 2>&1
	@cat ioid/artifacts/contracts/ioIDRegistry.sol/ioIDRegistry.json | jq .abi > abis/ioIDRegistry.json
	@echo DONE

.PHONY: generate_abi_project_device
generate_abi_project_device:
	@echo "#### generate abis of ProjectDevice from latest contracts..."
	@cd ../../smartcontracts && yarn install > /dev/null 2>&1 && yarn hardhat compile > /dev/null 2>&1
	@cat ../../smartcontracts/artifacts/contracts/examples/ProjectDevice.sol/ProjectDevice.json | jq .abi > abis/ProjectDevice.json
	@echo DONE

.PHONY: generate_go
generate_go:
	@echo "#### generate go code by abis"
	@go generate .
	@echo DONE

.PHONY: clean
clean:
	@rm -rf ioid

