package advent

// MaxInt returns the maximum value of two ints
func MaxInt(x, y int) int {
    if x < y {
        return y
    }

    return x
}
