package t

// T Ternary Operator
func T[A any](cond bool, a func() A, b func() A) A {
	if cond {
		return a()
	}
	return b()
}

// L aux function for T
func L[A any](a A) func() A {
	return func() A {
		return a
	}
}
