package example_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/ashashev/go-generics-example/example"
	"github.com/ashashev/go-generics-example/util/call"
	"github.com/ashashev/go-generics-example/util/gen"
	"github.com/ashashev/go-generics-example/util/slice"
)

func BenchmarkSliceContainsString(b *testing.B) {
	data := slice.Fill(30, call.Semi1(gen.RandStr, 20))
	needleExist := data[len(data)/2]
	needleNotExist := "unknown"

	b.Run("Generic Doesn't Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice.Contains(data, needleNotExist)
		}
	})

	b.Run("Generic Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice.Contains(data, needleExist)
		}
	})

	b.Run("Specific Doesn't Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			example.SliceContainsString(data, needleNotExist)
		}
	})

	b.Run("Specific Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			example.SliceContainsString(data, needleExist)
		}
	})
}

func BenchmarkSliceContainsInt(b *testing.B) {
	data := slice.Fill(30, call.Semi1(rand.Intn, math.MaxInt))
	needleExist := data[len(data)/2]
	needleNotExist := -5

	b.Run("Generic Doesn't Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice.Contains(data, needleNotExist)
		}
	})

	b.Run("Generic Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slice.Contains(data, needleExist)
		}
	})

	b.Run("Specific Doesn't Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			example.SliceContainsInt(data, needleNotExist)
		}
	})

	b.Run("Specific Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			example.SliceContainsInt(data, needleExist)
		}
	})
}
