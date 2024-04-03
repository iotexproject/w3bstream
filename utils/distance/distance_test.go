package distance

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPoint(t *testing.T) {
	r := require.New(t)

	point1 := MustNewPoint(uint64(0))
	t.Log(point1.Value())

	point2 := MustNewPoint(uint32(3))
	t.Log(point2.Value())

	r.Equal(point2.Distance(point1), point1.Distance(point2))

	defer func() {
		t.Log(recover())
	}()
	_ = MustNewPoint(int(1))
}

func TestNewPoints(t *testing.T) {
	r := require.New(t)

	values := []any{"1", "2", "3", uint64(4), []byte("5")}
	points := NewPoints(values...)
	r.Equal(points.Values(), values)

	distances := points.Distances(MustNewPoint("any"))
	for _, v := range distances {
		t.Log(v.Int64())
	}
}
