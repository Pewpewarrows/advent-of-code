package advent

import (
    "sort"
)

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

// MeanInts returns the mean value for a slice of ints
func MeanInts(xs []int) int {
    sort.Ints(xs)
    return xs[MedianInts(xs) - 1]
}

// MedianInts returns the median value for a slice of ints
func MedianInts(xs []int) int {
    return (len(xs) + 1) / 2
}
