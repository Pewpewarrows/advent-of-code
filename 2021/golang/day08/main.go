package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "strings"
)

func main() {
    var data []displayPattern
    advent.Execute(scanInputData, &data)

    count := uniqueSegmentedDigitCount(data)

    fmt.Println("solution:", count)
}

type displayPattern struct {
    signalPatterns [10]string
    displayOutput [4]string
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    data := *inputDataPtr.(*[]displayPattern)

    for scanner.Scan() {
        patterns := [10]string{}
        output := [4]string{}
        halves := strings.Split(scanner.Text(), "|")

        for i, text := range strings.Fields(halves[0]) {
            patterns[i] = text
        }

        for i, text := range strings.Fields(halves[1]) {
            output[i] = text
        }

        data = append(data, displayPattern{
            patterns,
            output,
        })
    }

    *inputDataPtr.(*[]displayPattern) = data
}

func uniqueSegmentedDigitCount(data []displayPattern) (count int) {
    sevenSegmentSignalsByNumber := map[int]int{
        0: 6,
        1: 2,
        2: 5,
        3: 5,
        4: 4,
        5: 5,
        6: 6,
        7: 3,
        8: 7,
        9: 6,
    }

    // TODO: compute this dynamically
    uniqueSegmentedDigits := []int{
        sevenSegmentSignalsByNumber[1],
        sevenSegmentSignalsByNumber[4],
        sevenSegmentSignalsByNumber[7],
        sevenSegmentSignalsByNumber[8],
    }

    for _, pattern := range data {
        for _, digit := range pattern.displayOutput {
            for _, i := range uniqueSegmentedDigits {
                if len(digit) == i {
                    count++
                    break
                }
            }
        }
    }

    return
}
