package distance

import (
	"crypto/sha256"
	"math/big"
	"sort"

	"github.com/machinefi/sprout/util/hash"
)

type hashDistance struct {
	distance *big.Int
	hash     [sha256.Size]byte
}

func Sort(locations []uint64, myLocation uint64) []uint64 {
	locationMap := map[[sha256.Size]byte]uint64{}
	for _, l := range locations {
		locationMap[hash.Sum256Uint64(l)] = l
	}
	myLocationHash := hash.Sum256Uint64(myLocation)
	ds := make([]hashDistance, 0, len(locations))

	for h := range locationMap {
		d := new(big.Int).Xor(new(big.Int).SetBytes(h[:]), new(big.Int).SetBytes(myLocationHash[:]))
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
