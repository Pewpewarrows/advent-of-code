package main

import (
    "fmt"
)

func main() {
}

// MaxInt returns the maximum value of two ints
func MaxInt(x, y int) int {
    if x < y {
        return y
    }

    return x
}

type subDirection int

const (
    undefined subDirection = iota
    forward
    down
    up
)

func (d subDirection) String() string {
    switch d {
    case forward:
        return "forward"
    case down:
        return "down"
    case up:
        return "up"
    }

    return "unknown"
}

type subCommand struct {
    direction subDirection
    magnitude int
}

func coordinatesFromCourse(course []subCommand) (hPos int, depth int) {
    for _, command := range course {
        switch command.direction {
        case forward:
            hPos += command.magnitude
        case down:
            depth += command.magnitude
        case up:
            depth -= command.magnitude
            // don't allow depth above the surface
            depth = MaxInt(0, depth)
        }
        fmt.Println(hPos, depth)
    }

    return
}
