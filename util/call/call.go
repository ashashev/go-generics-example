package call

func Semi1[A any, R any](f func(A) R, a A) func() R {
	return func() R {
		return f(a)
	}
}

func Semi2[T1, T2 any, R any](f func(T1, T2) R, v1 T1) func(T2) R {
	return func(v2 T2) R {
		return f(v1, v2)
	}
}

func Semi3[T1, T2, T3, R any](f func(T1, T2, T3) R, v1 T1) func(T2, T3) R {
	return func(v2 T2, v3 T3) R {
		return f(v1, v2, v3)
	}
}

func Semi4[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, v1 T1) func(T2, T3, T4) R {
	return func(v2 T2, v3 T3, v4 T4) R {
		return f(v1, v2, v3, v4)
	}
}

func Semi5[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) R, v1 T1) func(T2, T3, T4, T5) R {
	return func(v2 T2, v3 T3, v4 T4, v5 T5) R {
		return f(v1, v2, v3, v4, v5)
	}
}
