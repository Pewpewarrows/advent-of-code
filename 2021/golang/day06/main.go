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

    solution := simulatingLanternfishPopulation(fishSchool, 80)

    fmt.Println("solution:", solution)
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
