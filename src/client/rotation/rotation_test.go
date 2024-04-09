package rotation

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestFindRank(t *testing.T) {
	// goerli block number: 10000000
	blockHash := common.HexToHash("0x5977ca3297dee3788c0affe880011f86051d7f06196048bb938dc089e992bc56")
	watchTowerAddresses := [2]common.Address{
		common.HexToAddress("0xA6275EE214F80A532C3aBEe0A4cbBC2D1Dc22A72"),
		common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"),
	}

	watchTowerClientAddress := common.HexToAddress("0xA6275EE214F80A532C3aBEe0A4cbBC2D1Dc22A72")

	actualRank := findRank(blockHash, watchTowerAddresses[:], watchTowerClientAddress)

	if actualRank < 0 || actualRank >= len(watchTowerAddresses) {
		t.Fatalf("Rank %v is out of valid range [0, %v)", actualRank, len(watchTowerAddresses))
	}

	expectedRank := 0
	if actualRank != expectedRank {
		t.Fatalf("Rank Actual: %v, Expected: %v", actualRank, expectedRank)
	}
}

func TestAssignChain(t *testing.T) {
	hashRank := 1
	rollups := []Rollup{{1, "arbitrum", 1}, {2, "base", 1}}
	actual_chain := assignChain(hashRank, rollups)
	expected_chain := &Rollup{2, "base", 1}
	if *actual_chain != *expected_chain {
		t.Fatalf("Actual chain: %v != Expected chain: %v", actual_chain, expected_chain)
	}
}
