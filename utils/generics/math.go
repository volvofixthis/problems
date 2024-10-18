package generics

import "golang.org/x/exp/constraints"

type Num interface {
	constraints.Float | constraints.Signed
}

func Max[T Num](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T Num](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Abs[T Num](a T) T {
	if a < T(0) {
		a = T(-1) * a
	}

	return a
}

func AbsDiff[T Num](a, b T) T {
	c := a - b
	return Abs(c)
}
