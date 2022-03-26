package example_test

import (
	"testing"

	"github.com/ashashev/go-generics-example/example"
	"github.com/ashashev/go-generics-example/util/slice"
	"github.com/ashashev/go-generics-example/util/maps"
)

func BenchmarkGetKeys(b *testing.B) {
	data := slice.ToMap(
		slice.Fill(30, example.GenData),
		func(d *example.Data) string { return d.ID },
	)

	b.Run("Generic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maps.GetKeys(data)
		}
	})

	b.Run("Specific", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			example.GetKeysStringData(data)
		}
	})
}
