package result

import (
	"fmt"

	"github.com/ashashev/go-generics-example/data/either"
)

var ErrEmptyResult = fmt.Errorf("empty result")

type Result[T any] either.Either[error, T]

func FromOK[T any](value T) Result[T] {
	return either.FromRight[error](value)
}

func FromErr[T any](value error) Result[T] {
	return either.FromLeft[T](value)
}

func Map[T1, T2 any](r1 Result[T1], op func(T1)T2) Result[T2] {
	return either.Map[error, T1](r1, op)
}
