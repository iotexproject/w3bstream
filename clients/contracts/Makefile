.DEFAULT_GOAL := update

.PHONY: update
update: clean fetch generate_abi generate_go

.PHONY: fetch
fetch:
	@echo "#### clone ioID-contracts main branch..."
	@git clone --quiet git@github.com:machinefi/ioID-contracts.git ioid
	@echo DONE

.PHONY: generate_abi
generate_abi:
	@echo "#### generate abis from latest contracts..."
	@cd ioid && yarn install > /dev/null 2>&1 && yarn hardhat compile > /dev/null 2>&1
	@cd ioid/artifacts/contracts && for file in 'ioID' 'ioIDRegistry' ; \
	do \
		if [ -e $$file.sol/$$file.json ]; then \
			cat $$file.sol/$$file.json | jq .abi > ../../../abis/$$file.json; \
		fi \
	done
	@echo DONE

.PHONY: generate_go
generate_go:
	@echo "#### generate go code by abis"
	@go generate .
	@echo DONE

.PHONY: clean
clean:
	@rm -rf ioid

