package either

import (
	"fmt"

	"github.com/ashashev/go-generics-example/util/call"
)

var ErrEitherNoLeft = fmt.Errorf("either doesn't have left value")
var ErrEitherNoRight = fmt.Errorf("either doesn't have right value")
var ErrUnexpectedEitherBehaiviour = fmt.Errorf("unexpected result behaviour")

type Either[L, R any] interface {
	GetLeft() L
	GetRight() R
	IsLeft() bool
	IsRight() bool
	Swap() Either[R, L]

	restrictedEitherImplementation()
}

func FromLeft[R, L any](value L) *Left[L, R] {
	return &Left[L, R]{
		Value: value,
	}
}

func FromRight[L, R any](value R) *Right[L, R] {
	return &Right[L, R]{
		Value: value,
	}
}

func Map[L, A, B any](op func(A) B, a Either[L, A]) Either[L, B] {
	switch v := a.(type) {
	case Left[L, A]:
		return FromLeft[B](v.Value)
	case Right[L, A]:
		return FromRight[L](op(v.Value))
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func Map2[L, A1, A2, B any](op func(A1, A2) B, a1 Either[L, A1], a2 Either[L, A2]) Either[L, B] {
	switch v := a1.(type) {
	case Left[L, A1]:
		return FromLeft[B](v.Value)
	case Right[L, A1]:
		return Map(call.Semi2(op, v.Value), a2)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func Map3[L, A1, A2, A3, B any](op func(A1, A2, A3) B, a1 Either[L, A1], a2 Either[L, A2], a3 Either[L, A3]) Either[L, B] {
	switch v := a1.(type) {
	case Left[L, A1]:
		return FromLeft[B](v.Value)
	case Right[L, A1]:
		return Map2(call.Semi3(op, v.Value), a2, a3)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func Map4[L, A1, A2, A3, A4, B any](op func(A1, A2, A3, A4) B, a1 Either[L, A1], a2 Either[L, A2], a3 Either[L, A3], a4 Either[L, A4]) Either[L, B] {
	switch v := a1.(type) {
	case Left[L, A1]:
		return FromLeft[B](v.Value)
	case Right[L, A1]:
		return Map3(call.Semi4(op, v.Value), a2, a3, a4)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func Map5[L, A1, A2, A3, A4, A5, B any](op func(A1, A2, A3, A4, A5) B, a1 Either[L, A1], a2 Either[L, A2], a3 Either[L, A3], a4 Either[L, A4], a5 Either[L, A5]) Either[L, B] {
	switch v := a1.(type) {
	case Left[L, A1]:
		return FromLeft[B](v.Value)
	case Right[L, A1]:
		return Map4(call.Semi5(op, v.Value), a2, a3, a4, a5)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func FlatMap[L, A, B any](op func(A) Either[L, B], a Either[L, A]) Either[L, B] {
	switch v := a.(type) {
	case Left[L, A]:
		return FromLeft[B](v.Value)
	case Right[L, A]:
		return op(v.Value)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func FlatMap2[L, A1, A2, B any](op func(A1, A2) Either[L, B], a1 Either[L, A1], a2 Either[L, A2]) Either[L, B] {
	switch v := a1.(type) {
	case Left[L, A1]:
		return FromLeft[B](v.Value)
	case Right[L, A1]:
		return FlatMap(call.Semi2(op, v.Value), a2)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func FlatMap3[L, A1, A2, A3, B any](op func(A1, A2, A3) Either[L, B], a1 Either[L, A1], a2 Either[L, A2], a3 Either[L, A3]) Either[L, B] {
	switch v := a1.(type) {
	case Left[L, A1]:
		return FromLeft[B](v.Value)
	case Right[L, A1]:
		return FlatMap2(call.Semi3(op, v.Value), a2, a3)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func FlatMap4[L, A1, A2, A3, A4, B any](op func(A1, A2, A3, A4) Either[L, B], a1 Either[L, A1], a2 Either[L, A2], a3 Either[L, A3], a4 Either[L, A4]) Either[L, B] {
	switch v := a1.(type) {
	case Left[L, A1]:
		return FromLeft[B](v.Value)
	case Right[L, A1]:
		return FlatMap3(call.Semi4(op, v.Value), a2, a3, a4)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

func FlatMap5[L, A1, A2, A3, A4, A5, B any](op func(A1, A2, A3, A4, A5) Either[L, B], a1 Either[L, A1], a2 Either[L, A2], a3 Either[L, A3], a4 Either[L, A4], a5 Either[L, A5]) Either[L, B] {
	switch v := a1.(type) {
	case Left[L, A1]:
		return FromLeft[B](v.Value)
	case Right[L, A1]:
		return FlatMap4(call.Semi5(op, v.Value), a2, a3, a4, a5)
	}
	panic(ErrUnexpectedEitherBehaiviour)
}

type baseEitherImpl[L, R any] struct{}

func (baseEitherImpl[L, r]) GetLeft() L {
	panic(ErrEitherNoLeft)
}

func (baseEitherImpl[L, R]) GetRight() R {
	panic(ErrEitherNoRight)
}

func (baseEitherImpl[L, R]) restrictedEitherImplementation() {}

type Left[L, R any] struct {
	baseEitherImpl[L, R]
	Value L
}

func (l Left[L, R]) GetLeft() L {
	return l.Value
}

func (l Left[L, R]) Swap() Either[R, L] {
	return FromRight[R](l.Value)
}

func (Left[L, R]) IsLeft() bool {
	return true
}

func (Left[L, R]) IsRight() bool {
	return false
}

type Right[L, R any] struct {
	baseEitherImpl[L, R]
	Value R
}

func (r Right[L, R]) GetRight() R {
	return r.Value
}

func (r Right[L, R]) Swap() Either[R, L] {
	return FromLeft[L](r.Value)
}

func (Right[L, R]) IsLeft() bool {
	return false
}

func (Right[L, R]) IsRight() bool {
	return true
}
