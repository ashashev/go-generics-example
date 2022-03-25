package foo_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/ashashev/go-generics-example/foo"
	"github.com/ashashev/go-generics-example/gen"
	"github.com/ashashev/go-generics-example/util"
)

var (
	getKeysData = util.SliceToMap(
		util.FillSlice(30, foo.GenData),
		func(d *foo.Data) string { return d.ID },
	)

	sliceContainsStringData = util.FillSlice(30, util.SemiCall1(gen.RandStr, 20))
	needleStringExist       = sliceContainsStringData[len(sliceContainsStringData)/2]
	needleStringNotExist    = "unknown"

	sliceContainsIntData = util.FillSlice(30, util.SemiCall1(rand.Intn, math.MaxInt))
	needleIntExist       = sliceContainsIntData[len(sliceContainsIntData)/2]
	needleIntNotExist    = -5
)

func BenchmarkGetKeys(b *testing.B) {
	b.Run("Generic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.GetKeys(getKeysData)
		}
	})

	b.Run("Specific", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.GetKeysStringData(getKeysData)
		}
	})
}

func BenchmarkSliceContains(b *testing.B) {
	b.Run("String Generic Doesn't Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.SliceContains(sliceContainsStringData, needleStringNotExist)
		}
	})

	b.Run("String Generic Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.SliceContains(sliceContainsStringData, needleStringExist)
		}
	})

	b.Run("String Specific Doesn't Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.SliceContainsString(sliceContainsStringData, needleStringNotExist)
		}
	})

	b.Run("String Specific Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.SliceContainsString(sliceContainsStringData, needleStringExist)
		}
	})

	b.Run("Int Generic Doesn't Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.SliceContains(sliceContainsIntData, needleIntNotExist)
		}
	})

	b.Run("Int Generic Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.SliceContains(sliceContainsIntData, needleIntExist)
		}
	})

	b.Run("Int Specific Doesn't Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.SliceContainsInt(sliceContainsIntData, needleIntNotExist)
		}
	})

	b.Run("Int Specific Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			foo.SliceContainsInt(sliceContainsIntData, needleIntExist)
		}
	})
}
