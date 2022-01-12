package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "math"
    "sort"
    "strconv"
    "strings"
)

func main() {
    var crabYs []int
    advent.Execute(scanInputData, &crabYs)

    fuelCount := cheapestFuelForAlignment(crabYs)
    fmt.Println("part one:", fuelCount)

    fuelCount = cheapestIncrementalFuelForAlignment(crabYs)
    fmt.Println("part two:", fuelCount)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    crabYs := *inputDataPtr.(*[]int)

    for scanner.Scan() {
        for _, text := range strings.Split(scanner.Text(), ",") {
            i, err := strconv.ParseInt(text, 10, 32)
            if err != nil {
                // TODO: handle me
            }
            // TODO: why is i here an int64?
            crabYs = append(crabYs, int(i))
        }
    }

    *inputDataPtr.(*[]int) = crabYs
}

func cheapestFuelForAlignment(crabYs []int) (fuelCount int) {
    median := advent.MedianInts(crabYs)

    for _, y := range crabYs {
        fuelCount += int(math.Abs(float64(y) - median))
    }

    return
}

func cheapestIncrementalFuelForAlignment(crabYs []int) (fuelCount int) {
    // brute force search
    cheapestFuelCount := 0
    cheapestAlignment := 0
    sort.Ints(crabYs)

    for i := crabYs[0]; i < crabYs[len(crabYs) - 1]; i++ {
        fuelCount = 0

        for _, y := range crabYs {
            fuelCount += advent.TriangleNumberInt(int(math.Abs(float64(y - i))))
        }

        if (cheapestFuelCount == 0) || (fuelCount < cheapestFuelCount) {
            cheapestFuelCount = fuelCount
            cheapestAlignment = i
        }
    }

    // TODO: only while debug
    fmt.Println("cheapestAlignment:", cheapestAlignment)

    fuelCount = cheapestFuelCount

    return
}
