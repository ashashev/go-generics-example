package result1

import (
	"fmt"

	"github.com/ashashev/go-generics-example/util/call"
)

var ErrEmptyResult = fmt.Errorf("empty result")
var ErrUnexpectedResultBehaiviour = fmt.Errorf("unexpected result behaviour")

type Result[T any] interface {
	Get() T
	Err() error

	// Both returns default value for T if result is empty
	// Both() (T, error)

	restrictedImplementation()
}

func NewOK[T any](value T) Result[T] {
	return &OK[T]{value: value}
}

func NewErr[T any](err error) Result[T] {
	return &Err[T]{err: err}
}

func Map[A, B any](r Result[A], op func(A) B) Result[B] {
	switch v := r.(type) {
	case *OK[A]:
		return &OK[B]{value: op(v.value)}
	case *Err[A]:
		return &Err[B]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func Map2[T1, T2, R any](r1 Result[T1], r2 Result[T2], op func(T1, T2) R) Result[R] {
	switch v := r1.(type) {
	case *OK[T1]:
		return Map(r2, call.Semi2(op, v.value))
	case *Err[T1]:
		return &Err[R]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func Map3[T1, T2, T3, R any](r1 Result[T1], r2 Result[T2], r3 Result[T3], op func(T1, T2, T3) R) Result[R] {
	switch v := r1.(type) {
	case *OK[T1]:
		return Map2(r2, r3, call.Semi3(op, v.value))
	case *Err[T1]:
		return &Err[R]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func Map4[T1, T2, T3, T4, R any](r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], op func(T1, T2, T3, T4) R) Result[R] {
	switch v := r1.(type) {
	case *OK[T1]:
		return Map3(r2, r3, r4, call.Semi4(op, v.value))
	case *Err[T1]:
		return &Err[R]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func Map5[T1, T2, T3, T4, T5, R any](r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5], op func(T1, T2, T3, T4, T5) R) Result[R] {
	switch v := r1.(type) {
	case *OK[T1]:
		return Map4(r2, r3, r4, r5, call.Semi5(op, v.value))
	case *Err[T1]:
		return &Err[R]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap[A, B any](r Result[A], op func(A) Result[B]) Result[B] {
	switch v := r.(type) {
	case *OK[A]:
		return op(v.value)
	case *Err[A]:
		return &Err[B]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap2[T1, T2, R any](r1 Result[T1], r2 Result[T2], op func(T1, T2) Result[R]) Result[R] {
	switch v := r1.(type) {
	case *OK[T1]:
		return FlatMap(r2, call.Semi2(op, v.value))
	case *Err[T1]:
		return &Err[R]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap3[T1, T2, T3, R any](r1 Result[T1], r2 Result[T2], r3 Result[T3], op func(T1, T2, T3) Result[R]) Result[R] {
	switch v := r1.(type) {
	case *OK[T1]:
		return FlatMap2(r2, r3, call.Semi3(op, v.value))
	case *Err[T1]:
		return &Err[R]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap4[T1, T2, T3, T4, R any](r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], op func(T1, T2, T3, T4) Result[R]) Result[R] {
	switch v := r1.(type) {
	case *OK[T1]:
		return FlatMap3(r2, r3, r4, call.Semi4(op, v.value))
	case *Err[T1]:
		return &Err[R]{err: v.err}
	}
	panic(ErrUnexpectedResultBehaiviour)
}

func FlatMap5[T1, T2, T3, T4, T5, R any](r1 Result[T1], r2 Result[T2], r3 Result[T3], r4 Result[T4], r5 Result[T5], op func(T1, T2, T3, T4, T5) Result[R]) Result[R] {
	switch v := r1.(type) {
	case *OK[T1]:
		return FlatMap4(r2, r3, r4, r5, call.Semi5(op, v.value))
	case *Err[T1]:
		return &Err[R]{err: v.err}
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
	value T
}

func (r *OK[T]) Get() T {
	return r.value
}

type Err[T any] struct {
	baseResultImpl[T]
	err error
}

func (r *Err[T]) Err() error {
	return r.err
}
