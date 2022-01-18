package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    var grid octopusGrid
    advent.Execute(scanInputData, &grid)
    grid2 := octopusGrid{}
    grid.Copy(&grid2)

    flashCount := flashCountSimulation(grid, 100)
    fmt.Println("part one:", flashCount)

    steps := stepsUntilFlashSync(grid2)
    fmt.Println("part two:", steps)
}

type octopusGrid struct {
    energyLevels [][]int
    stepCount int
    flashCount int
}

func (o *octopusGrid) String() string {
    return fmt.Sprintf("step #%d with %d cumulative flashes", o.stepCount, o.flashCount)
}

// in production code instead use a lib like https://github.com/jinzhu/copier
func (o *octopusGrid) Copy(dupe *octopusGrid) {
    dupe.energyLevels = make([][]int, len(o.energyLevels))
    // copy(dupe.energyLevels, o.energyLevels)
    for i := range dupe.energyLevels {
        dupe.energyLevels[i] = make([]int, len(o.energyLevels[i]))
        copy(dupe.energyLevels[i], o.energyLevels[i])
    }

    dupe.stepCount = o.stepCount
    dupe.flashCount = o.flashCount
}

func (o *octopusGrid) step() {
    flashGrid := make([][]bool, len(o.energyLevels))
    for i := range flashGrid {
        flashGrid[i] = make([]bool, len(o.energyLevels[i]))
    }

    for i := 0; i < len(o.energyLevels); i++ {
        for j := 0; j < len(o.energyLevels[i]); j++ {
            o.energyLevels[i][j]++
        }
    }

    for {
        didFlash := false

        for i := 0; i < len(o.energyLevels); i++ {
            hasUp := true
            if i == 0 {
                hasUp = false
            }

            hasDown := true
            if i == (len(o.energyLevels) - 1) {
                hasDown = false
            }

            for j := 0; j < len(o.energyLevels[i]); j++ {
                hasLeft := true
                if j == 0 {
                    hasLeft = false
                }

                hasRight := true
                if j == (len(o.energyLevels[i]) - 1) {
                    hasRight = false
                }

                if o.energyLevels[i][j] < 10 {
                    continue
                }

                if flashGrid[i][j] {
                    continue
                }

                flashGrid[i][j] = true
                didFlash = true
                o.flashCount++

                if hasUp {
                    if hasLeft {
                        o.energyLevels[i - 1][j - 1]++
                    }
                    o.energyLevels[i - 1][j]++
                    if hasRight {
                        o.energyLevels[i - 1][j + 1]++
                    }
                }
                if hasRight {
                    o.energyLevels[i][j + 1]++
                }
                if hasDown {
                    if hasRight {
                        o.energyLevels[i + 1][j + 1]++
                    }
                    o.energyLevels[i + 1][j]++
                    if hasLeft {
                        o.energyLevels[i + 1][j - 1]++
                    }
                }
                if hasLeft {
                    o.energyLevels[i][j - 1]++
                }
            }
        }

        if !didFlash {
            break
        }
    }

    for i := 0; i < len(o.energyLevels); i++ {
        for j := 0; j < len(o.energyLevels[i]); j++ {
            if (o.energyLevels[i][j] > 9) {
                o.energyLevels[i][j] = 0
            }
        }
    }

    o.stepCount++
}

func (o *octopusGrid) flashIsSynced() bool {
    prevEnergyLevel := o.energyLevels[0][0]

    for i := 0; i < len(o.energyLevels); i++ {
        for j := 0; j < len(o.energyLevels[i]); j++ {
            if (o.energyLevels[i][j] != prevEnergyLevel) {
                return false
            }

            prevEnergyLevel = o.energyLevels[i][j]
        }
    }

    return true
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    grid := *inputDataPtr.(*octopusGrid)
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
        grid.energyLevels = append(grid.energyLevels, row)
    }

    *inputDataPtr.(*octopusGrid) = grid
}

func flashCountSimulation(grid octopusGrid, stepCount int) (flashCount int) {
    for s := 0; s < stepCount; s++ {
        grid.step()
    }

    return grid.flashCount
}

func stepsUntilFlashSync(grid octopusGrid) (steps int) {
    for {
        grid.step()

        if grid.flashIsSynced() {
            return grid.stepCount
        }
    }
}
