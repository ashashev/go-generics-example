package example_test

import (
	"testing"

	"github.com/ashashev/go-generics-example/data/result"
	"github.com/ashashev/go-generics-example/example"
)

func DataBuildedWithIfGuard(firstFail bool) (*example.Data, error) {
	s := &example.S1{}

	parts, err := s.GetPartsID(firstFail)
	if err != nil {
		return nil, err
	}

	id, err := s.JoinParts(parts)
	if err != nil {
		return nil, err
	}

	field1, err := s.GetField1(parts, id)
	if err != nil {
		return nil, err
	}

	field2, err := s.GetField2(parts, field1)
	if err != nil {
		return nil, err
	}

	field3, err := s.GetField3(field1, field2)
	if err != nil {
		return nil, err
	}

	return example.NewData(id, field1, field2, field3), nil
}

func DataBuildedWithChain(firstFail bool) (*example.Data, error) {
	dc := &example.DataConstrutor{
		S1:   &example.S1{},
		Fail: firstFail,
	}

	err := dc.Chain(
		dc.GetPartsID,
		dc.JoinParts,
		dc.GetField1,
		dc.GetField2,
		dc.GetField3,
	)
	if err != nil {
		return nil, err
	}

	return dc.GetData(), nil
}

func DataBuildedWithResult(firstFail bool) (*example.Data, error) {
	s := &example.S2{}

	parts := s.GetPartsID(firstFail)
	id := result.FlatMap(parts, s.JoinParts)
	field1 := result.FlatMap2(parts, id, s.GetField1)
	field2 := result.FlatMap2(parts, field1, s.GetField2)
	field3 := result.FlatMap2(field1, field2, s.GetField3)

	return result.Map4(id, field1, field2, field3, example.NewData).Both()
}

func BenchmarkResultSuccess(b *testing.B) {
	b.Run("IfGuard", func(b *testing.B) {
		DataBuildedWithIfGuard(false)
	})

	b.Run("Chain", func(b *testing.B) {
		DataBuildedWithChain(false)
	})

	b.Run("Result", func(b *testing.B) {
		DataBuildedWithResult(false)
	})
}

func BenchmarkResultFirstFail(b *testing.B) {
	b.Run("IfGuard", func(b *testing.B) {
		DataBuildedWithIfGuard(true)
	})

	b.Run("Chain", func(b *testing.B) {
		DataBuildedWithChain(true)
	})

	b.Run("Result", func(b *testing.B) {
		DataBuildedWithResult(true)
	})
}
