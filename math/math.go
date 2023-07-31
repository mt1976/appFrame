package math

// Max returns the larger of x or y.
// The max function returns the larger of two integers.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
// The function "min" returns the smaller of two integers.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
