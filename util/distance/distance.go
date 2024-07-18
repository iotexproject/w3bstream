package distance

import (
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"

	"github.com/iotexproject/w3bstream/util/hash"
)

type hashDistance struct {
	distance *big.Int
	hash     common.Hash
}

func Sort(locations []uint64, myLocation uint64) []uint64 {
	locationMap := map[common.Hash]uint64{}
	for _, l := range locations {
		locationMap[hash.Keccak256Uint64(l)] = l
	}
	myLocationHash := hash.Keccak256Uint64(myLocation)
	ds := make([]hashDistance, 0, len(locations))

	for h := range locationMap {
		d := new(big.Int).Xor(h.Big(), myLocationHash.Big())
		ds = append(ds, hashDistance{
			distance: d,
			hash:     h,
		})
	}

	sort.Slice(ds, func(i, j int) bool {
		return ds[i].distance.Cmp(ds[j].distance) < 0
	})

	result := []uint64{}
	for _, d := range ds {
		result = append(result, locationMap[d.hash])
	}
	return result
}
