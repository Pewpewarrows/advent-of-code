package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
)

func main() {
    var octopusGrid [10][10]int
    advent.Execute(scanInputData, &octopusGrid)

    flashCount := flashCountSimulation(octopusGrid, 100)

    fmt.Println("solution:", flashCount)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    octopusGrid := *inputDataPtr.(*[]int)
    var i int

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d", &i)
        octopusGrid = append(octopusGrid, i)
    }

    *inputDataPtr.(*[]int) = octopusGrid
}

func flashCountSimulation(octopusGrid [10][10]int, stepCount int) (flashCount int) {
    fmt.Println(octopusGrid)
    for s := 0; s < stepCount; s++ {
        var nextOctopusGrid [10][10]int

        // first compute new energy levels
        for i := 0; i < len(octopusGrid); i++ {
            hasUp := true
            if i == 0 {
                hasUp = false
            }

            hasDown := true
            if i == (len(octopusGrid) - 1) {
                hasDown = false
            }

            for j := 0; j < len(octopusGrid[i]); j++ {
                hasLeft := true
                if j == 0 {
                    hasLeft = false
                }

                hasRight := true
                if j == (len(octopusGrid[i]) - 1) {
                    hasRight = false
                }

                // count the surrounding 9+s on all eight adjacent positions
                adjacentFlashCount := 0
                if hasUp {
                    if hasLeft {
                        if octopusGrid[i - 1][j - 1] >= 9 {
                            adjacentFlashCount++
                        }
                    }
                    if octopusGrid[i - 1][j] >= 9 {
                        adjacentFlashCount++
                    }
                    if hasRight {
                        if octopusGrid[i - 1][j + 1] >= 9 {
                            adjacentFlashCount++
                        }
                    }
                }
                if hasRight {
                    if octopusGrid[i][j + 1] >= 9 {
                        adjacentFlashCount++
                    }
                }
                if hasDown {
                    if hasRight {
                        if octopusGrid[i + 1][j + 1] >= 9 {
                            adjacentFlashCount++
                        }
                    }
                    if octopusGrid[i + 1][j] >= 9 {
                        adjacentFlashCount++
                    }
                    if hasLeft {
                        if octopusGrid[i + 1][j - 1] >= 9 {
                            adjacentFlashCount++
                        }
                    }
                }
                if hasLeft {
                    if octopusGrid[i][j - 1] >= 9 {
                        adjacentFlashCount++
                    }
                }

                nextOctopusGrid[i][j] = (octopusGrid[i][j] + 1 + adjacentFlashCount)
            }
        }

        // TODO: this isn't adequate, may be a chain reaction of flashes
        octopusGrid = nextOctopusGrid

        // then count flashes and reset to 0
        for i := 0; i < len(octopusGrid); i++ {
            for j := 0; j < len(octopusGrid[i]); j++ {
                if (octopusGrid[i][j] > 9) {
                    flashCount++
                    octopusGrid[i][j] = 0
                }
            }
        }

        if (s == 0) || (s == 1) || (s == 2) {
            fmt.Println(octopusGrid)
        }
    }

    return
}
