package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "sort"
    "strconv"
    "strings"
)

func main() {
    var heightMap [][]int
    advent.Execute(scanInputData, &heightMap)

    sum := summedLowPointRiskLevels(heightMap)
    fmt.Println("part one:", sum)

    product := productOfLargestBasins(heightMap, 3)
    fmt.Println("part two:", product)
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

func productOfLargestBasins(heightMap [][]int, basinCount int) (product int) {
    var basinSizes []int
    var visitedCoords [][2]int

    for i := range heightMap {
        for j := range heightMap[i] {
            if coordWasVisited([2]int{i, j}, visitedCoords) {
                continue
            }

            if heightMap[i][j] == 9 {
                continue
            }

            visitedCoords = append(visitedCoords, [2]int{i, j})
            coords := adjacentBasinCoords(heightMap, i, j)
            basinSize := 1

            for {
                var newCoords [][2]int
                for _, coord := range coords {
                    if coordWasVisited(coord, visitedCoords) {
                        continue
                    }

                    newCoords = append(newCoords, adjacentBasinCoords(heightMap, coord[0], coord[1])...)
                    visitedCoords = append(visitedCoords, coord)
                    basinSize++
                }
                coords = newCoords

                if len(coords) == 0 {
                    break
                }
            }

            basinSizes = append(basinSizes, basinSize)
        }
    }

    sort.Ints(basinSizes)

    product = 1
    for i := (len(basinSizes) - 1); i > (len(basinSizes) - 1 - basinCount); i-- {
        product *= basinSizes[i]
    }

    return
}

func coordWasVisited(coord [2]int, visitedCoords [][2]int) bool {
    for _, v := range visitedCoords {
        if v == coord {
            return true
        }
    }

    return false
}

func adjacentBasinCoords(heightMap [][]int, i, j int) (adjacentCoords [][2]int) {
    width := len(heightMap[0])
    height := len(heightMap)

    if i != 0 {
        height := heightMap[i - 1][j]
        if height != 9 {
            adjacentCoords = append(adjacentCoords, [2]int{i - 1, j})
        }
    }

    if i != (height - 1) {
        height := heightMap[i + 1][j]
        if height != 9 {
            adjacentCoords = append(adjacentCoords, [2]int{i + 1, j})
        }
    }

    if j != 0 {
        height := heightMap[i][j - 1]
        if height != 9 {
            adjacentCoords = append(adjacentCoords, [2]int{i, j - 1})
        }
    }

    if j != (width - 1) {
        height := heightMap[i][j + 1]
        if height != 9 {
            adjacentCoords = append(adjacentCoords, [2]int{i, j + 1})
        }
    }

    return
}
