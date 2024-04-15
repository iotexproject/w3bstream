package distance

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNearestN(t *testing.T) {
	r := require.New(t)

	from := uint32(1)
	to := []any{"1", "2", uint32(3)}
	n := 5

	values := NearestN(from, n, to...)
	r.Equal(values, to)

	n = 2
	values = NearestN(from, 2, to...)
	r.Len(values, n)
	t.Log(values)

	r.True(slices.Contains(values, "2"))
}
