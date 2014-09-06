package cycle_test

import (
	"testing"

	"github.com/levicook/cycle"
	"github.com/stretchr/testify/require"
)

var empty = cycle.Strings()
var simpsons = cycle.Strings("Marge", "Homer", "Bart", "Lisa", "Maggie")

func Benchmark_Strings_with_simpsons(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = simpsons.Next()
	}
}

func Test_Strings_with_empty(t *testing.T) {
	for i := 0; i < 10; i++ {
		require.Equal(t, "", empty.Next())
	}
}

func Test_Strings_with_simpsons(t *testing.T) {
	for i := 0; i < 10; i++ {
		require.Equal(t, "Marge", simpsons.Next())
		require.Equal(t, "Homer", simpsons.Next())
		require.Equal(t, "Bart", simpsons.Next())
		require.Equal(t, "Lisa", simpsons.Next())
		require.Equal(t, "Maggie", simpsons.Next())
	}
}
