package advent

// MaxInt returns the maximum value of two ints
func MaxInt(x, y int) int {
    if x < y {
        return y
    }

    return x
}

// MinInt returns the minimum value of two ints
func MinInt(x, y int) int {
    if x < y {
        return x
    }

    return y
}

// DivCeilInt returns the ceiled divison of two ints
func DivCeilInt(x, y int) int {
	return 1 + ((x - 1) / y)
}
