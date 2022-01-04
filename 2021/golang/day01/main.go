package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
)

func main() {
    livePtr := flag.Bool("live", false, "whether to fetch the input live from the internet")
    // TODO: usage instructions for args: single path to input file
    // TODO: also allow input to come via stdin, to use curl pipe
    flag.Parse()
    live := *livePtr
    filePaths := flag.Args()

    if live {
        if len(filePaths) > 0 {
            log.Fatalf("invalid args, may only fetch input live or supply file arg, not both: %s", filePaths)
        }
    } else {
        if len(filePaths) == 0 {
            log.Fatal("invalid args, expected one input file path, found none")
        } else if len(filePaths) > 1 {
            log.Fatalf("invalid args, expected only one input file path, found multiple: %s", filePaths)
        }
    }

    if live {
        // TODO: write me
        return
    }

    f, err := os.Open(filePaths[0])
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    depths := make([]int, 0)
    var i int
    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d", &i)
        depths = append(depths, i)
    }
	if err := scanner.Err(); err != nil {
        log.Fatal("could not read input:", err)
	}

    count := depthIncreaseCount(depths)
    slidingCount := slidingDepthIncreaseCount(depths)

    fmt.Println("part one:", count)
    fmt.Println("part two:", slidingCount)
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
    window := make([]int, windowLen)
    var windowDepth int
    var prevWindowDepth int

    for i, depth := range depths {
        if len(window) == windowLen {
            _, window = window[0], window[1:]
        }

        window = append(window, depth)

        // need a full window first
        if i < (windowLen - 1) {
            continue
        }

        windowDepth = 0
        for _, d := range window {
            windowDepth += d
        }

        // first window doesn't count as an increase
        if i == (windowLen - 1) {
            goto loop
        }

        if windowDepth > prevWindowDepth {
            count++
        }

loop:
        prevWindowDepth = windowDepth
    }

    return
}
