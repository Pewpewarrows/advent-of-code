package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

const birthTimer = 7
const startingBirthTimer = 8

func main() {
    var fishSchool []int
    advent.Execute(scanInputData, &fishSchool)

    // TODO: really have to figure out why the sim func mutates fishSchool
    // solution := simulatingLanternfishPopulation(fishSchool, 80)
    solution := estimatingLanternfishPopulation(fishSchool, 80)
    fmt.Println("part one:", solution)

    solution = estimatingLanternfishPopulation(fishSchool, 256)
    fmt.Println("part two:", solution)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    fishSchool := *inputDataPtr.(*[]int)

    for scanner.Scan() {
        for _, text := range strings.Split(scanner.Text(), ",") {
            i, err := strconv.ParseInt(text, 10, 32)
            if err != nil {
                // TODO: handle me
            }
            // TODO: why is i here an int64?
            fishSchool = append(fishSchool, int(i))
        }
    }

    *inputDataPtr.(*[]int) = fishSchool
}

func simulatingLanternfishPopulation(fishSchool []int, dayCount int) (fishCount int) {
    for day := 0; day < dayCount; day++ {
        addedFishCount := 0

        for i := range fishSchool {
            if fishSchool[i] == 0 {
                fishSchool[i] = birthTimer
                addedFishCount++
            }

            fishSchool[i]--
        }

        for i := 0; i < addedFishCount; i++ {
            fishSchool = append(fishSchool, startingBirthTimer)
        }
    }

    fishCount = len(fishSchool)

    return
}

// TODO: consider making fishCount be a big type?
func estimatingLanternfishPopulation(fishSchool []int, dayCount int) (fishCount int) {
    // this closed formula attempt didn't pan out...
    // fmt.Println(dayCount % birthTimer)
    // startingTimer := fishSchool[0]
    // birthCycleCount := (dayCount - startingTimer - 1) / birthTimer
    // if (dayCount > startingTimer) {
    //     // account for the first birth
    //     birthCycleCount++
    // }

    fishCountsByTimer := make(map[int]int)
    var largestTimer int

    for _, fish := range fishSchool {
        fishCountsByTimer[fish]++

        if fish > largestTimer {
            largestTimer = fish
        }
    }

    largestTimer = advent.MaxInt(largestTimer, startingBirthTimer)

    for day := 0; day < dayCount; day++ {
        pendingNewFish := fishCountsByTimer[0]
        for i := 1; i <= largestTimer + 1; i++ {
            fishCountsByTimer[i - 1] = fishCountsByTimer[i]
        }
        fishCountsByTimer[birthTimer - 1] += pendingNewFish
        fishCountsByTimer[startingBirthTimer] += pendingNewFish
    }

    for _, f := range fishCountsByTimer {
        fishCount += f
    }

    return
}
