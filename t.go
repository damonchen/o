package t

// L Ternary Operator for lazy evaluate
func L[A any](cond bool, a func() A, b func() A) A {
	if cond {
		return a()
	}
	return b()
}

// T Ternary Operator
func T[A any](cond bool, a A, b A) A {
	if cond {
		return a
	}
	return b
}

// W aux function for T
func W[A any](a A) func() A {
	return func() A {
		return a
	}
}
