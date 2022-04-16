package result1

import (
	"fmt"

	"github.com/ashashev/go-generics-example/util/call"
)

var ErrEmptyResult = fmt.Errorf("empty result")

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
	if !r.IsOK() {
		return Err[B]{Value: r.Err()}
	}

	return OK[B]{Value: op(r.Get())}
}

func Map2[T1, T2, R any](op func(T1, T2) R, r1 Result[T1], r2 Result[T2]) Result[R] {
	if !r1.IsOK() {
		return Err[R]{Value: r1.Err()}
	}

	return Map(call.Semi2(op, r1.Get()), r2)
}

func Map3[T1, T2, T3, R any](op func(T1, T2, T3) R, r1 Result[T1], r2 Result[T2], r3 Result[T3]) Result[R] {
	if !r1.IsOK() {
		return Err[R]{Value: r1.Err()}
	}

	return Map2(call.Semi3(op, r1.Get()), r2, r3)
}

func Map4[T1, T2, T3, T4, R any](op func(T1, T2, T3, T4) R, r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4]) Result[R] {
	if !r1.IsOK() {
		return Err[R]{Value: r1.Err()}
	}

	return Map3(call.Semi4(op, r1.Get()), r2, r3, r4)
}

func Map5[T1, T2, T3, T4, T5, R any](op func(T1, T2, T3, T4, T5) R, r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5]) Result[R] {
	if !r1.IsOK() {
		return Err[R]{Value: r1.Err()}
	}

	return Map4(call.Semi5(op, r1.Get()), r2, r3, r4, r5)
}

func FlatMap[A, B any](op func(A) Result[B], r Result[A]) Result[B] {
	if !r.IsOK() {
		return Err[B]{Value: r.Err()}
	}

	return op(r.Get())
}

func FlatMap2[T1, T2, R any](op func(T1, T2) Result[R], r1 Result[T1], r2 Result[T2]) Result[R] {
	if !r1.IsOK() {
		return Err[R]{Value: r1.Err()}
	}

	return FlatMap(call.Semi2(op, r1.Get()), r2)
}

func FlatMap3[T1, T2, T3, R any](op func(T1, T2, T3) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3]) Result[R] {
	if !r1.IsOK() {
		return Err[R]{Value: r1.Err()}
	}

	return FlatMap2(call.Semi3(op, r1.Get()), r2, r3)
}

func FlatMap4[T1, T2, T3, T4, R any](op func(T1, T2, T3, T4) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4]) Result[R] {
	if !r1.IsOK() {
		return Err[R]{Value: r1.Err()}
	}

	return FlatMap3(call.Semi4(op, r1.Get()), r2, r3, r4)
}

func FlatMap5[T1, T2, T3, T4, T5, R any](op func(T1, T2, T3, T4, T5) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5]) Result[R] {
	if !r1.IsOK() {
		return Err[R]{Value: r1.Err()}
	}

	return FlatMap4(call.Semi5(op, r1.Get()), r2, r3, r4, r5)
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
