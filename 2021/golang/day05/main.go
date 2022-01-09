package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
)

func main() {
    var ventLines []ventLine
    advent.Execute(scanInputData, &ventLines)

    solution := ventOverlapCount(ventLines)

    fmt.Println("solution:", solution)
}

// TODO: point struct
type ventLine struct {
    x1, y1, x2, y2 int
}

func (v ventLine) coords() (xys [][2]int) {
    // TODO: this only handles horizontal and vertical lines
    var vertical bool
    var smallerCoord, largerCoord int

    if v.x1 == v.x2 {
        vertical = true
        if v.y1 > v.y2 {
            largerCoord = v.y1
            smallerCoord = v.y2
        } else {
            largerCoord = v.y2
            smallerCoord = v.y1
        }
    } else if v.y1 == v.y2 {
        vertical = false
        if v.x1 > v.x2 {
            largerCoord = v.x1
            smallerCoord = v.x2
        } else {
            largerCoord = v.x2
            smallerCoord = v.x1
        }
    } else {
        return
    }

    xys = append(xys, [2]int{v.x1, v.y1})
    xys = append(xys, [2]int{v.x2, v.y2})

    for i := smallerCoord + 1; i < largerCoord; i++ {
        if vertical {
            xys = append(xys, [2]int{v.x1, i})
        } else {
            xys = append(xys, [2]int{i, v.y1})
        }
    }

    return
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    ventLines := *inputDataPtr.(*[]ventLine)
    var x1, y1, x2, y2 int

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
        ventLines = append(ventLines, ventLine{
            x1, y1, x2, y2,
        })
    }

    *inputDataPtr.(*[]ventLine) = ventLines
}

func ventOverlapCount(ventLines []ventLine) (overlapCount int) {
    var largestX, largestY int

    for _, line := range ventLines {
        if line.x1 > largestX {
            largestX = line.x1
        }

        if line.x2 > largestX {
            largestX = line.x2
        }

        if line.y1 > largestY {
            largestY = line.y1
        }

        if line.y2 > largestY {
            largestY = line.y2
        }
    }

    // assumes 0 is a valid coord
    floorGrid := make([][]int, largestX + 1)

    for x := range floorGrid {
        floorGrid[x] = make([]int, largestY + 1)
    }

    for _, line := range ventLines {
        for _, coord := range line.coords() {
            floorGrid[coord[0]][coord[1]]++
        }
    }

    for _, row := range floorGrid {
        for _, point := range row {
            if point > 1 {
                overlapCount++
            }
        }
    }

    return
}
