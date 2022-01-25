package main

import (
    "testing"
)

func TestFoldableGridDotCount(t *testing.T) {
    grid := newFoldableGrid(
        []coord{
            {6, 10},
            {0, 14},
            {9, 10},
            {0, 3},
            {10, 4},
            {4, 11},
            {6, 0},
            {6, 12},
            {4, 1},
            {0, 13},
            {10, 12},
            {3, 4},
            {3, 0},
            {8, 4},
            {1, 10},
            {2, 14},
            {8, 10},
            {9, 0},
        },
    )
    actual := grid.fold(fold{vertical, 7}).dotCount

    if actual != 17 {
        t.Errorf("example: expected 17 actual %d", actual)
    }
}
