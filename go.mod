module github.com/witnesschain-com/diligencewatchtower-client

go 1.22.1

toolchain go1.22.2

require (
	github.com/Layr-Labs/eigensdk-go v0.0.8
	github.com/ethereum/go-ethereum v1.14.3
	github.com/gorilla/websocket v1.5.1
	github.com/jellydator/ttlcache/v3 v3.1.1
	github.com/urfave/cli/v2 v2.25.7
	github.com/witnesschain-com/diligencewatchtower-client/common v0.0.0-00010101000000-000000000000
	github.com/witnesschain-com/diligencewatchtower-client/external v0.0.0-00010101000000-000000000000
	github.com/witnesschain-com/diligencewatchtower-client/opchain v0.0.0-00010101000000-000000000000
	github.com/witnesschain-com/diligencewatchtower-client/watcher v0.0.0-00010101000000-000000000000
	github.com/witnesschain-com/diligencewatchtower-client/webserver v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.24.0
)

require (
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.10.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/cbergoon/merkletree v0.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/crate-crypto/go-kzg-4844 v1.0.0 // indirect
	github.com/deckarep/golang-set/v2 v2.1.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/howeyc/gopass v0.0.0-20210920133722-c8aef6fb66ef // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/supranational/blst v0.3.11 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/wagslane/go-password-validator v0.3.0 // indirect
	github.com/witnesschain-com/diligencewatchtower-client/L1 v0.0.0-00010101000000-000000000000 // indirect
	github.com/witnesschain-com/diligencewatchtower-client/bindings v0.0.0-00010101000000-000000000000 // indirect
	github.com/witnesschain-com/diligencewatchtower-client/contractutils v0.0.0-00010101000000-000000000000 // indirect
	github.com/witnesschain-com/operator-cli v0.1.1-0.20240620133912-3518d3c5bff1 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/term v0.19.0 // indirect
	golang.org/x/tools v0.20.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)

replace github.com/witnesschain-com/diligencewatchtower-client/L1 => ./L1

replace github.com/witnesschain-com/diligencewatchtower-client/common => ./common

replace github.com/witnesschain-com/diligencewatchtower-client/contractutils => ./contractutils

replace github.com/witnesschain-com/diligencewatchtower-client/bindings => ./bindings

replace github.com/witnesschain-com/diligencewatchtower-client/webserver => ./webserver

replace github.com/witnesschain-com/diligencewatchtower-client/opchain => ./L2/opchain

replace github.com/witnesschain-com/diligencewatchtower-client/rotation => ./rotation

replace github.com/witnesschain-com/diligencewatchtower-client/watcher => ./watcher

replace github.com/witnesschain-com/diligencewatchtower-client/external => ./external
