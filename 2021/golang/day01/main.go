package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
)

func main() {
    var depths []int
    advent.Execute(scanInputData, &depths)

    count := depthIncreaseCount(depths)
    slidingCount := slidingDepthIncreaseCount(depths)

    fmt.Println("part one:", count)
    fmt.Println("part two:", slidingCount)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    depths := *inputDataPtr.(*[]int)
    var i int

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d", &i)
        depths = append(depths, i)
    }

    *inputDataPtr.(*[]int) = depths
}

func depthIncreaseCount(depths []int) (count int) {
    var prevDepth int

    for i, depth := range depths {
        // first depth doesn't count as an increase
        if i == 0 {
            goto loop
        }

        if depth > prevDepth {
            count++
        }

loop:
        prevDepth = depth
    }

    return
}

func slidingDepthIncreaseCount(depths []int) (count int) {
    windowLen := 3

    for i, depth := range depths {
        if i < windowLen {
            continue
        }

        if depths[i - 3] < depth {
            count++
        }
    }

    return
}
