RPC="https://blue-orangutan-rpc.eu-north-2.gateway.fm/"

test: test-watcher test-opchain test-contractutils
	@echo "test completes"

test-watcher:
	go test github.com/witnesschain-com/diligencewatchtower-client/watcher

test-opchain:
	go test github.com/witnesschain-com/diligencewatchtower-client/opchain

test-contractutils-anvil:
	docker run --rm --detach --name op_sepolia --network=host ghcr.io/foundry-rs/foundry:latest 'anvil --fork-url ${RPC} '
	@sleep 15
	-go test github.com/witnesschain-com/diligencewatchtower-client/contractutils
	docker stop op_sepolia

test-contractutils:
	go test github.com/witnesschain-com/diligencewatchtower-client/contractutils

anvil: 
	docker run --rm --name op_sepolia --network=host ghcr.io/foundry-rs/foundry:latest 'anvil --fork-url ${RPC}'
