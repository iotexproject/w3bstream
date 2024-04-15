package distance

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log/slog"
	"math/big"
	"sort"

	"github.com/machinefi/sprout/util/hash"
)

type hashDistance struct {
	distance *big.Int
	hash     [sha256.Size]byte
}

func GetMinNLocation(locations []uint64, myLocation, n uint64) []uint64 {
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
	ds = ds[:n]
	for _, d := range ds {
		result = append(result, locationMap[d.hash])
	}
	return result
}

func NewPoint(val any) (*Point, error) {
	var (
		buf = bytes.NewBuffer([]byte{})
		err error
	)

	switch v := val.(type) {
	case string:
		_, err = buf.WriteString(v)
	case []byte:
		_, err = buf.Write(v)
	default:
		err = binary.Write(buf, binary.LittleEndian, v)
	}
	if err != nil {
		return nil, err
	}
	h := sha256.Sum256(buf.Bytes())
	return &Point{
		value: val,
		point: new(big.Int).SetBytes(h[:]),
	}, nil
}

func MustNewPoint(val any) *Point {
	point, err := NewPoint(val)
	if err != nil {
		slog.Error("failed to new pointer", "value", val, "error", err)
		panic(err)
	}
	return point
}

type Point struct {
	value any
	point *big.Int
}

func (p *Point) Distance(other *Point) *big.Int {
	return new(big.Int).Xor(p.point, other.point)
}

func (p *Point) Value() any {
	return p.value
}

func NewPoints(values ...any) Points {
	points := make(Points, 0, len(values))
	for i := range values {
		points = append(points, MustNewPoint(values[i]))
	}
	return points
}

type Points []*Point

func (ps Points) Values() []any {
	values := make([]any, 0, len(ps))
	for _, p := range ps {
		values = append(values, p.value)
	}
	return values
}

func (ps Points) Distances(point *Point) []*big.Int {
	distances := make([]*big.Int, 0, len(ps))
	for _, p := range ps {
		distances = append(distances, p.Distance(point))
	}
	return distances
}
