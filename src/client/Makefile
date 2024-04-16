L1RPC="ws://52.90.204.64:8547/"
L2RPC="http://35.153.33.57:8545/"

test: test-watcher test-opchain test-contractutils
	@echo "test completes"

test-rotation: 
	docker run --rm --detach --name foundry  ghcr.io/foundry-rs/foundry:latest anvil --broadcast --fork-url ${L1RPC} 
	-go test -v github.com/witnesschain-com/diligencewatchtower-client/watchtower/rotation
	docker stop foundry

test-watcher:
	go test github.com/witnesschain-com/diligencewatchtower-client/watcher

test-opchain:
	go test github.com/witnesschain-com/diligencewatchtower-client/opchain

test-contractutils:
	docker run --rm --detach --name op_sepolia --network=host ghcr.io/foundry-rs/foundry:latest 'anvil --chain-id 11155111 --fork-url ${L1RPC} '
	@sleep 15
	-go test github.com/witnesschain-com/diligencewatchtower-client/contractutils
	docker stop op_sepolia
