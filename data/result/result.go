package result

import (
	"fmt"

	"github.com/ashashev/go-generics-example/util/call"
)

var ErrEmptyResult = fmt.Errorf("empty result")
var ErrUnexpectedResultBehaiviour = fmt.Errorf("unexpected result behaviour")

type Result[T any] interface {
	Get() T
	Err() error
	IsOK() bool

	// Both returns default value for T if result is empty
	Both() (T, error)

	restrictedImplementation()
}

func FromOK[T any](value T) Result[T] {
	return OK[T]{Value: value}
}

func FromErr[T any](err error) Result[T] {
	return Err[T]{Value: err}
}

func FromBoth[T any](value T, err error) Result[T] {
	if err != nil {
		return FromErr[T](err)
	}

	return FromOK(value)
}

func Map[A, B any](op func(A) B, r Result[A]) Result[B] {
	switch v := r.(type) {
	case OK[A]:
		return OK[B]{Value: op(v.Value)}
	case Err[A]:
		return Err[B]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func Map2[T1, T2, R any](op func(T1, T2) R, r1 Result[T1], r2 Result[T2]) Result[R] {
	switch v := r1.(type) {
	case OK[T1]:
		return Map(call.Semi2(op, v.Value), r2)
	case Err[T1]:
		return Err[R]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func Map3[T1, T2, T3, R any](op func(T1, T2, T3) R, r1 Result[T1], r2 Result[T2], r3 Result[T3]) Result[R] {
	switch v := r1.(type) {
	case OK[T1]:
		return Map2(call.Semi3(op, v.Value), r2, r3)
	case Err[T1]:
		return Err[R]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func Map4[T1, T2, T3, T4, R any](op func(T1, T2, T3, T4) R, r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4]) Result[R] {
	switch v := r1.(type) {
	case OK[T1]:
		return Map3(call.Semi4(op, v.Value), r2, r3, r4)
	case Err[T1]:
		return Err[R]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func Map5[T1, T2, T3, T4, T5, R any](op func(T1, T2, T3, T4, T5) R, r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5]) Result[R] {
	switch v := r1.(type) {
	case OK[T1]:
		return Map4(call.Semi5(op, v.Value), r2, r3, r4, r5)
	case Err[T1]:
		return Err[R]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap[A, B any](op func(A) Result[B], r Result[A]) Result[B] {
	switch v := r.(type) {
	case OK[A]:
		return op(v.Value)
	case Err[A]:
		return Err[B]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap2[T1, T2, R any](op func(T1, T2) Result[R], r1 Result[T1], r2 Result[T2]) Result[R] {
	switch v := r1.(type) {
	case OK[T1]:
		return FlatMap(call.Semi2(op, v.Value), r2)
	case Err[T1]:
		return Err[R]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap3[T1, T2, T3, R any](op func(T1, T2, T3) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3]) Result[R] {
	switch v := r1.(type) {
	case OK[T1]:
		return FlatMap2(call.Semi3(op, v.Value), r2, r3)
	case Err[T1]:
		return Err[R]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap4[T1, T2, T3, T4, R any](op func(T1, T2, T3, T4) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4]) Result[R] {
	switch v := r1.(type) {
	case OK[T1]:
		return FlatMap3(call.Semi4(op, v.Value), r2, r3, r4)
	case Err[T1]:
		return Err[R]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap5[T1, T2, T3, T4, T5, R any](op func(T1, T2, T3, T4, T5) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5]) Result[R] {
	switch v := r1.(type) {
	case OK[T1]:
		return FlatMap4(call.Semi5(op, v.Value), r2, r3, r4, r5)
	case Err[T1]:
		return Err[R]{Value: v.Value}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

type baseResultImpl[T any] struct{}

func (baseResultImpl[T]) Get() T {
	panic(ErrEmptyResult)
}

func (baseResultImpl[T]) Err() error {
	return nil
}

func (baseResultImpl[T]) restrictedImplementation() {}

type OK[T any] struct {
	baseResultImpl[T]
	Value T
}

func (r OK[T]) Get() T {
	return r.Value
}

func (r OK[T]) Both() (T, error) {
	return r.Value, nil
}

func (OK[T]) IsOK() bool {
	return true
}

type Err[T any] struct {
	baseResultImpl[T]
	Value error
}

func (r Err[T]) Err() error {
	return r.Value
}

func (r Err[T]) Both() (value T, err error) {
	err = r.Value
	return
}

func (Err[T]) IsOK() bool {
	return false
}
