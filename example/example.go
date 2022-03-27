package example

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"

	"github.com/ashashev/go-generics-example/data/result"
	"github.com/ashashev/go-generics-example/util/gen"
)

type Data struct {
	ID     string
	Field1 string
	Field2 float64
	Field3 uint64
}

func NewData(ID string, Field1 string, Field2 float64, Field3 uint64) *Data {
	return &Data{
		ID:     ID,
		Field1: Field1,
		Field2: Field2,
		Field3: Field3,
	}
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

type DataConstrutor struct {
	S1     *S1
	Fail   bool
	parts  []string
	id     string
	field1 string
	field2 float64
	field3 uint64
}

func (dc *DataConstrutor) GetPartsID() (err error) {
	dc.parts, err = dc.S1.GetPartsID(dc.Fail)
	return
}

func (dc *DataConstrutor) JoinParts() (err error) {
	dc.id, err = dc.S1.JoinParts(dc.parts)
	return
}

func (dc *DataConstrutor) GetField1() (err error) {
	dc.field1, err = dc.S1.GetField1(dc.parts, dc.id)
	return
}

func (dc *DataConstrutor) GetField2() (err error) {
	dc.field2, err = dc.S1.GetField2(dc.parts, dc.field1)
	return
}

func (dc *DataConstrutor) GetField3() (err error) {
	dc.field3, err = dc.S1.GetField3(dc.field1, dc.field2)
	return
}

func (dc *DataConstrutor) GetData() *Data {
	return NewData(dc.id, dc.field1, dc.field2, dc.field3)
}

func (*DataConstrutor) Chain(ops ...func() error) error {
	for _, op := range ops {
		if err := op(); err != nil {
			return err
		}
	}
	return nil
}

type S1 struct{}

func (*S1) GetPartsID(fail bool) ([]string, error) {
	if fail {
		return nil, fmt.Errorf("fail")
	}

	return []string{"1qaz", "2wsx"}, nil
}

func (*S1) JoinParts(ps []string) (string, error) {
	return "1qaz-2wsx", nil
}

func (*S1) GetField1([]string, string) (string, error) {
	return "3edc", nil
}

func (*S1) GetField2([]string, string) (float64, error) {
	return 0.34, nil
}

func (*S1) GetField3(string, float64) (uint64, error) {
	return 24, nil
}

type S2 struct{}

func (*S2) GetPartsID(fail bool) result.Result[[]string] {
	if fail {
		return result.FromErr[[]string](fmt.Errorf("fail"))
	}

	return result.FromOK([]string{"1qaz", "2wsx"})
}

func (*S2) JoinParts(ps []string) result.Result[string] {
	return result.FromOK("1qaz-2wsx")
}

func (*S2) GetField1([]string, string) result.Result[string] {
	return result.FromOK("3edc")
}

func (*S2) GetField2([]string, string) result.Result[float64] {
	return result.FromOK(0.34)
}

func (*S2) GetField3(string, float64) result.Result[uint64] {
	return result.FromOK[uint64](24)
}
