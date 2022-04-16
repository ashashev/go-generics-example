package result2

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

type result [T any] struct {
	value T
	err error
}

func FromOK[T any](value T) Result[T] {
	return &result[T]{value: value}
}

func FromErr[T any](err error) Result[T] {
	return &result[T]{err: err}
}

func FromBoth[T any](value T, err error) Result[T] {
	if err != nil {
		return &result[T]{err: err}
	}

	return FromOK(value)
}

func Map[A, B any](op func(A) B, r Result[A]) Result[B] {
	if !r.IsOK() {
		return &result[B]{err: r.Err()}
	}

	return &result[B]{value: op(r.Get())}
}

func Map2[T1, T2, R any](op func(T1, T2) R, r1 Result[T1], r2 Result[T2]) Result[R] {
	if !r1.IsOK() {
		return &result[R]{err: r1.Err()}
	}

	return Map(call.Semi2(op, r1.Get()), r2)
}

func Map3[T1, T2, T3, R any](op func(T1, T2, T3) R, r1 Result[T1], r2 Result[T2], r3 Result[T3]) Result[R] {
	if !r1.IsOK() {
		return &result[R]{err: r1.Err()}
	}

	return Map2(call.Semi3(op, r1.Get()), r2, r3)
}

func Map4[T1, T2, T3, T4, R any](op func(T1, T2, T3, T4) R, r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4]) Result[R] {
	if !r1.IsOK() {
		return &result[R]{err: r1.Err()}
	}

	return Map3(call.Semi4(op, r1.Get()), r2, r3, r4)
}

func Map5[T1, T2, T3, T4, T5, R any](op func(T1, T2, T3, T4, T5) R, r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5]) Result[R] {
	if !r1.IsOK() {
		return &result[R]{err: r1.Err()}
	}

	return Map4(call.Semi5(op, r1.Get()), r2, r3, r4, r5)
}

func FlatMap[A, B any](op func(A) Result[B], r Result[A]) Result[B] {
	if !r.IsOK() {
		return &result[B]{err: r.Err()}
	}

	return op(r.Get())
}

func FlatMap2[T1, T2, R any](op func(T1, T2) Result[R], r1 Result[T1], r2 Result[T2]) Result[R] {
	if !r1.IsOK() {
		return &result[R]{err: r1.Err()}
	}

	return FlatMap(call.Semi2(op, r1.Get()), r2)
}

func FlatMap3[T1, T2, T3, R any](op func(T1, T2, T3) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3]) Result[R] {
	if !r1.IsOK() {
		return &result[R]{err: r1.Err()}
	}

	return FlatMap2(call.Semi3(op, r1.Get()), r2, r3)
}

func FlatMap4[T1, T2, T3, T4, R any](op func(T1, T2, T3, T4) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4]) Result[R] {
	if !r1.IsOK() {
		return &result[R]{err: r1.Err()}
	}

	return FlatMap3(call.Semi4(op, r1.Get()), r2, r3, r4)
}

func FlatMap5[T1, T2, T3, T4, T5, R any](op func(T1, T2, T3, T4, T5) Result[R], r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5]) Result[R] {
	if !r1.IsOK() {
		return &result[R]{err: r1.Err()}
	}

	return FlatMap4(call.Semi5(op, r1.Get()), r2, r3, r4, r5)
}

func (r *result[T]) Get() T {
	if r.err != nil {
		panic(ErrEmptyResult)
	}

	return r.value
}

func (r *result[T]) Err() error {
	return r.err
}

func (r *result[T]) Both() (T, error) {
	return r.value, r.err
}

func (r *result[T]) IsOK() bool {
	return r.err == nil
}

func (result[T]) restrictedImplementation() {}
