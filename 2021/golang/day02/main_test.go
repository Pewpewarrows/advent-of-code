package main

import (
    "testing"
)

func TestCoordinatesFromCourse(t *testing.T) {
    course := []subCommand{
        {forward, 1},
        {down, 1},
    }
    hPos, depth := coordinatesFromCourse(course)

    if (hPos != 1) || (depth != 1) {
        t.Errorf("simple: expected (1, 1) actual (%d, %d)", hPos, depth)
    }

    course = []subCommand{
        {forward, 1},
        {forward, 2},
        {down, 1},
        {forward, 3},
        {up, 2},
        {down, 3},
        {down, 4},
    }
    hPos, depth = coordinatesFromCourse(course)

    if (hPos != 6) || (depth != 7) {
        t.Errorf("complex: expected (6, 7) actual (%d, %d)", hPos, depth)
    }

    course = []subCommand{}
    hPos, depth = coordinatesFromCourse(course)

    if (hPos != 0) || (depth != 0) {
        t.Errorf("empty: expected (0, 0) actual (%d, %d)", hPos, depth)
    }
}
