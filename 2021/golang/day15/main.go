package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "math"
)

func main() {
    var grid [][]int
    advent.Execute(scanInputData, &grid)

    solution := dijkstraShortestPath(grid, coord{0, 0}, coord{len(grid) - 1, len(grid[len(grid) - 1]) - 1})

    fmt.Println("solution:", solution)
}

type coord struct {
    x, y int
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    grid := *inputDataPtr.(*[]int)
    var i int

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d", &i)
        grid = append(grid, i)
    }

    *inputDataPtr.(*[]int) = grid
}

func dijkstraShortestPath(grid [][]int, source coord, target coord) (totalRisk int) {
    visitedGrid := make([][]bool, len(grid))
    for i := range visitedGrid {
        visitedGrid[i] = make([]bool, len(grid[i]))
    }

    dist := make(map[coord]float64)
    prev := make(map[coord]coord)

    for i, row := range grid {
        for j := range row {
            dist[coord{i, j}] = math.Inf(1)
            // prev[coord{i, j}] = nil
        }
    }
    dist[coord{0, 0}] = 0

    for i, row := range grid {
        for j, risk := range row {
            // if visitedGrid[i][j] {
            //     continue
            // }

            if ((i - 1) >= 0) && !visitedGrid[i - 1][j] {
                visitedGrid[i - 1][j] = true
            }

            if ((j + 1) < len(row)) && !visitedGrid[i][j + 1] {
                visitedGrid[i][j + 1] = true
            }

            if ((i + 1) < len(grid)) && !visitedGrid[i + 1][j] {
                visitedGrid[i + 1][j] = true
            }

            if ((j - 1) >= 0) && !visitedGrid[i][j - 1] {
                visitedGrid[i][j - 1] = true
            }
        }
    }

    return
}
