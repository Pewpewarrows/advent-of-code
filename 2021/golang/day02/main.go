package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
)

func main() {
    var course []subCommand
    advent.Execute(inputFromScanner, &course)

    hPos, depth := coordinatesFromCourse(course)

    fmt.Println("horizontal position:", hPos)
    fmt.Println("depth:", depth)
    fmt.Println("solution:", hPos * depth)
}

func inputFromScanner(scanner *bufio.Scanner, inputDataPtr interface{}) {
    course := *inputDataPtr.(*[]subCommand)
    var direction string
    var magnitude int

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%s %d", &direction, &magnitude)
        course = append(course, subCommand{
            directionFromString(direction),
            magnitude,
        })
    }

    *inputDataPtr.(*[]subCommand) = course
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

// TODO: naming convention for conversion func like this?
func directionFromString(s string) subDirection {
    switch s {
    case "forward":
        return forward
    case "down":
        return down
    case "up":
        return up
    }

    return undefined
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
            depth = advent.MaxInt(0, depth)
        }
    }

    return
}
