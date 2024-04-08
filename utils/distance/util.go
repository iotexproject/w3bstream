package distance

import (
	"math/big"
	"sort"
)

func NearestN(from any, n int, to ...any) []any {
	if len(to) <= n {
		return to
	}

	point := MustNewPoint(from)
	points := NewPoints(to...)

	distances := make([]*struct {
		distance *big.Int
		index    int
	}, 0, len(points))
	for i, p := range points {
		distances = append(distances, &struct {
			distance *big.Int
			index    int
		}{
			distance: p.Distance(point),
			index:    i,
		})
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance.Cmp(distances[j].distance) < 0
	})

	values := make([]any, 0, n)
	for i := 0; i < n; i++ {
		values = append(values, points[distances[i].index].Value())
	}
	return values
}
