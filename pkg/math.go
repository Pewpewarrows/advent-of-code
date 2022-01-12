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
func MeanInts(xs []int) (avg float64) {
    // TODO: 0-len check

    for _, x := range xs {
        avg += float64(x)
    }

    avg /= float64(len(xs))

    return
}

// MedianInts returns the median value for a slice of ints
func MedianInts(xs []int) float64 {
    // TODO: 0-len check

    sort.Ints(xs)

    // even
    if (len(xs) % 2) == 0 {
        return MeanInts([]int{xs[(len(xs) / 2) - 1], xs[len(xs) / 2]})
    }

    // odd
    i := (len(xs) + 1) / 2
    return float64(xs[i - 1])
}

// TriangleNumberInt returns the sum of all numbers 1...n
func TriangleNumberInt(n int) int {
    return (n * (n + 1)) / 2
}
