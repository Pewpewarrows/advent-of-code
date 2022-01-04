package main

import (
    "testing"
)

func TestDepthIncreaseCount(t *testing.T) {
    depths := []int{1, 2, 3}
    actual := depthIncreaseCount(depths)

    if 2 != actual {
        t.Errorf("simple: expected 2 actual %d", actual)
    }

    depths = []int{0, 3, 2, 1, 4}
    actual = depthIncreaseCount(depths)

    if 2 != actual {
        t.Errorf("complex: expected 2 actual %d", actual)
    }

    depths = []int{}
    actual = depthIncreaseCount(depths)

    if 0 != actual {
        t.Errorf("empty: expected 0 actual %d", actual)
    }

    depths = []int{42}
    actual = depthIncreaseCount(depths)

    if 0 != actual {
        t.Errorf("single: expected 0 actual %d", actual)
    }

    depths = []int{3, 2, 1}
    actual = depthIncreaseCount(depths)

    if 0 != actual {
        t.Errorf("only decreasing: expected 0 actual %d", actual)
    }
}
