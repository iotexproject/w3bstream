package distance

import (
	"crypto/sha256"
	"math/big"
	"sort"

	"github.com/machinefi/sprout/utils/hash"
)

type hashDistance struct {
	distance *big.Int
	hash     [sha256.Size]byte
}

func GetMinNLocation(locations []string, myLocation, n uint64) []string {
	locationMap := map[[sha256.Size]byte]string{}
	for _, n := range locations {
		locationMap[sha256.Sum256([]byte(n))] = n
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

	result := []string{}
	ds = ds[:n]
	for _, d := range ds {
		result = append(result, locationMap[d.hash])
	}
	return result
}
