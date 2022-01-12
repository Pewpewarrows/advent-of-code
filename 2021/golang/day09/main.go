package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    var heightMap [][]int
    advent.Execute(scanInputData, &heightMap)

    solution := summedLowPointRiskLevels(heightMap)

    fmt.Println("solution:", solution)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    heightMap := *inputDataPtr.(*[][]int)
    var line string

    for scanner.Scan() {
        var row []int
        fmt.Sscanf(scanner.Text(), "%s", &line)
        for _, s := range strings.Split(line, "") {
            i, err := strconv.ParseInt(s, 10, 32)
            if err != nil {
                // TODO: handle me
            }
            // TODO: why is i here an int64?
            row = append(row, int(i))
        }
        heightMap = append(heightMap, row)
    }

    *inputDataPtr.(*[][]int) = heightMap
}

func summedLowPointRiskLevels(heightMap [][]int) (sum int) {
    lowPointCoords := make([][2]int, 0)
    width := len(heightMap[0])
    height := len(heightMap)

    for i := range heightMap {
        for j := range heightMap[i] {
            var adjacentHeights []int

            if i != 0 {
                adjacentHeights = append(adjacentHeights, heightMap[i - 1][j])
            }

            if i != (height - 1) {
                adjacentHeights = append(adjacentHeights, heightMap[i + 1][j])
            }

            if j != 0 {
                adjacentHeights = append(adjacentHeights, heightMap[i][j - 1])
            }

            if j != (width - 1) {
                adjacentHeights = append(adjacentHeights, heightMap[i][j + 1])
            }

            isLowPoint := true
            for _, h := range adjacentHeights {
                if h <= heightMap[i][j] {
                    isLowPoint = false
                    break
                }
            }

            if isLowPoint {
                lowPointCoords = append(lowPointCoords, [2]int{i, j})
            }
        }
    }

    for _, coords := range lowPointCoords {
        sum += (heightMap[coords[0]][coords[1]] + 1)
    }

    return
}
