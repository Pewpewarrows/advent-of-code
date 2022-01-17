package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    var octopusGrid [10][10]int
    advent.Execute(scanInputData, &octopusGrid)

    flashCount := flashCountSimulation(octopusGrid, 100)

    fmt.Println("solution:", flashCount)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    octopusGrid := *inputDataPtr.(*[10][10]int)
    lineNum := 0
    var line string

    for scanner.Scan() {
        var row [10]int
        fmt.Sscanf(scanner.Text(), "%s", &line)
        for n, s := range strings.Split(line, "") {
            i, err := strconv.ParseInt(s, 10, 32)
            if err != nil {
                // TODO: handle me
            }
            // TODO: why is i here an int64?
            row[n] = int(i)
        }
        octopusGrid[lineNum] = row
        lineNum++
    }

    *inputDataPtr.(*[10][10]int) = octopusGrid
}

func flashCountSimulation(octopusGrid [10][10]int, stepCount int) (flashCount int) {
    for s := 0; s < stepCount; s++ {
        var flashGrid [10][10]bool

        for i := 0; i < len(octopusGrid); i++ {
            for j := 0; j < len(octopusGrid[i]); j++ {
                octopusGrid[i][j]++
            }
        }

        for {
            didFlash := false

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

                    if octopusGrid[i][j] < 10 {
                        continue
                    }

                    if flashGrid[i][j] {
                        continue
                    }

                    flashGrid[i][j] = true
                    didFlash = true
                    flashCount++

                    if hasUp {
                        if hasLeft {
                            octopusGrid[i - 1][j - 1]++
                        }
                        octopusGrid[i - 1][j]++
                        if hasRight {
                            octopusGrid[i - 1][j + 1]++
                        }
                    }
                    if hasRight {
                        octopusGrid[i][j + 1]++
                    }
                    if hasDown {
                        if hasRight {
                            octopusGrid[i + 1][j + 1]++
                        }
                        octopusGrid[i + 1][j]++
                        if hasLeft {
                            octopusGrid[i + 1][j - 1]++
                        }
                    }
                    if hasLeft {
                        octopusGrid[i][j - 1]++
                    }
                }
            }

            if !didFlash {
                break
            }
        }

        for i := 0; i < len(octopusGrid); i++ {
            for j := 0; j < len(octopusGrid[i]); j++ {
                if (octopusGrid[i][j] > 9) {
                    octopusGrid[i][j] = 0
                }
            }
        }
    }

    return
}
