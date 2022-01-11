package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "math"
    "strconv"
    "strings"
)

func main() {
    var crabYs []int
    advent.Execute(scanInputData, &crabYs)

    fuelCount := cheapestFuelForAlignment(crabYs)

    fmt.Println("solution:", fuelCount)
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
    mean := advent.MeanInts(crabYs)

    for _, y := range crabYs {
        fuelCount += int(math.Abs(float64(y - mean)))
    }

    return
}
