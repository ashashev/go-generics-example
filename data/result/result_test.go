package result_test

import (
	"fmt"
	"testing"

	"github.com/ashashev/go-generics-example/data/result"
	"github.com/ashashev/go-generics-example/data/result1"
	"github.com/ashashev/go-generics-example/data/result2"
)

func BenchmarkMapOK(b *testing.B) {
	src := result.FromOK("")
	op := func(string) int {return 0}
	src1 := result1.FromOK("")
	src2 := result2.FromOK("")


	b.Run("result Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result.Map(op, src)
		}
	})

	b.Run("result1 Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result1.Map(op, src1)
		}
	})

	b.Run("result2 Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result2.Map(op, src2)
		}
	})
}

func BenchmarkMapErr(b *testing.B) {
	src := result.FromErr[string](fmt.Errorf("error"))
	op := func(string) int {return 0}
	src1 := result1.FromErr[string](fmt.Errorf("error"))
	src2 := result2.FromErr[string](fmt.Errorf("error"))

	b.Run("result Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result.Map(op, src)
		}
	})

	b.Run("result1 Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result1.Map(op, src1)
		}
	})

	b.Run("result2 Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result2.Map(op, src2)
		}
	})
}
