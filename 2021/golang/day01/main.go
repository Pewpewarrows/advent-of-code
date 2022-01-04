package main

import (
    "fmt"
)

func main() {
    fmt.Println("letz heck thes")
}

func depthIncreaseCount(depths []int) (count int) {
    var prevDepth int

    for i, depth := range depths {
        // first depth doesn't count as an increase
        if i == 0 {
            prevDepth = depth
            continue
        }

        if depth > prevDepth {
            count++
        }

        prevDepth = depth
    }

    return
}
