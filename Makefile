LOCATION?="contract"
CONTRACT_LOCATION_DL?="https://raw.githubusercontent.com/uport-project/ethr-did-registry/develop/contracts/EthereumDIDRegistry.sol"
CONTRACT_LOCATION?=$(LOCATION)/EthereumDIDRegistry.sol
CONTRACT_BUILD?=$(LOCATION)/build
CONTRACT_ABI=$(CONTRACT_BUILD)/EthereumDIDRegistry.abi
CONTRACT_BIN=$(CONTRACT_BUILD)/EthereumDIDRegistry.bin
CONTRACT_GO?=$(LOCATION)/ether_did_registry.go

dir:
	mkdir -p $(LOCATION) || true

clean:
	rm -rf $(LOCATION) || true
	rm -rf $(CONTRACT_BUILD) || true

dl: dir
	curl $(CONTRACT_LOCATION_DL) -o $(CONTRACT_LOCATION)

abi:
	solc --overwrite --bin --abi $(CONTRACT_LOCATION) -o $(CONTRACT_BUILD)
	abigen --bin=$(CONTRACT_BIN) --abi=$(CONTRACT_ABI) --pkg=$(LOCATION) --out=$(CONTRACT_GO)