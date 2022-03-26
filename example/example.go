package example

import (
	"math/rand"

	"github.com/google/uuid"

	"github.com/ashashev/go-generics-example/util/gen"
)

type Data struct {
	ID     string
	Field1 string
	Field2 float64
	Field3 uint64
}

func GenData() *Data {
	return &Data{
		ID:     uuid.NewString(),
		Field1: gen.RandStr(rand.Intn(20)),
		Field2: rand.ExpFloat64(),
		Field3: rand.Uint64(),
	}
}

func GetKeysStringData(m map[string]*Data) []string {
	data := make([]string, 0, len(m))
	for k := range m {
		data = append(data, k)
	}
	return data
}

func SliceContainsString(haystack []string, needle string) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}

func SliceContainsInt(haystack []int, needle int) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}
