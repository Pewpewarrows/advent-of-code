package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
)

func main() {
    var p puzzle
    advent.Execute(scanInputData, &p)

    solution := p.fg.fold(p.folds[0]).dotCount

    fmt.Println("solution:", solution)
}

type puzzle struct {
    fg foldableGrid
    folds []fold
}

type fold struct {
    direction foldDirection
    // TODO: technically shouldn't the index always be centered based on dir?
    index int
}

type foldableGrid struct {
    grid [][]bool
    dotCount int
}

func newFoldableGrid(coords []coord) *foldableGrid {
    maxX := 0
    maxY := 0

    for _, c := range coords {
        if c.x > maxX {
            maxX = c.x
        }

        if c.y > maxY {
            maxY = c.y
        }
    }

    grid := make([][]bool, maxY + 1)
    for i := range grid {
        grid[i] = make([]bool, maxX + 1)
    }

    dotCount := 0

    for _, c := range coords {
        grid[c.y][c.x] = true
        dotCount++
    }

    return &foldableGrid{grid, dotCount}
}

func (g *foldableGrid) fold(f fold) foldableGrid {
    // TODO: validate that there are no dots on the fold line
    // TODO: validate that f.index is in-bounds and in middle

    grid := make([][]bool, len(g.grid))
    for i := range grid {
        grid[i] = make([]bool, len(g.grid[i]))
    }

    dotCount := 0

    for i, row := range g.grid {
        for j, isDot := range row {
            if !isDot {
                continue
            }

            // TODO: dedup these better
            if f.direction == vertical {
                if i < f.index {
                    grid[i][j] = true
                    dotCount++
                } else if i > f.index {
                    mirrorI := (f.index - (i - f.index))
                    if grid[mirrorI][j] {
                        continue
                    }

                    grid[mirrorI][j] = true
                    dotCount++
                }
            } else if f.direction == horizontal {
                if j < f.index {
                    grid[i][j] = true
                    dotCount++
                } else if j > f.index {
                    mirrorJ := (f.index - (j - f.index))
                    if grid[i][mirrorJ] {
                        continue
                    }

                    grid[i][mirrorJ] = true
                    dotCount++
                }
            } else {
                // TODO: error
            }
        }
    }

    return foldableGrid{grid, dotCount}
}

type foldDirection int

const (
    undefined foldDirection = iota
    vertical
    horizontal
)

// TODO: naming convention for conversion func like this?
func directionFromRune(r rune) foldDirection {
    switch r {
    case 'y':
        return vertical
    case 'x':
        return horizontal
    }

    return undefined
}

type coord struct {
    x int
    y int
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    p := *inputDataPtr.(*puzzle)
    var x, y int
    var direction rune
    var index int
    var coords []coord

    // TODO: state enum
    areCoordsExhausted := false

    for scanner.Scan() {
        if scanner.Text() == "" {
            areCoordsExhausted = true
            continue
        }

        if areCoordsExhausted {
            fmt.Sscanf(scanner.Text(), "fold along %c=%d", &direction, &index)
            p.folds = append(p.folds, fold{directionFromRune(direction), index})
        } else {
            fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
            coords = append(coords, coord{x, y})
        }
    }

    p.fg = *newFoldableGrid(coords)

    *inputDataPtr.(*puzzle) = p
}
