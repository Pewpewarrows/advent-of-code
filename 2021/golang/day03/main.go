package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "log"
    "math"
    "strconv"
)

func main() {
    var diagnostics []int
    advent.Execute(scanInputData, &diagnostics)

    gamma, epsilon := powerConsumptionFromDiagnostics(diagnostics)

    fmt.Println("gamma:", gamma)
    fmt.Println("epsilon:", epsilon)
    fmt.Println("part one:", gamma * epsilon)

    oxygen, co2 := lifeSupportFromDiagnostics(diagnostics)

    fmt.Println("oxygen:", oxygen)
    fmt.Println("co2:", co2)
    fmt.Println("part two:", oxygen * co2)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    diagnostics := *inputDataPtr.(*[]int)
    var s string

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%s", &s)
        code, err := strconv.ParseInt(s, 2, 32)
        if err != nil {
            // TODO: allow scanInputData to return errs
            log.Println("invalid int as binary string:", s, err)
        }
        diagnostics = append(diagnostics, int(code))
    }

    *inputDataPtr.(*[]int) = diagnostics
}

func bitCountsFromDiagnostics(diagnostics []int) (bitCounts map[int]int) {
    bitCounts = make(map[int]int)

    for _, code := range diagnostics {
        msb := int(math.Log2(float64(code)))

        for i := 0; i <= msb; i++ {
            bitCounts[i] += ((int(code) >> i) & 1)
        }
    }

    return
}

func positionBitCountFromDiagnostics(diagnostics []int, pos int) (bitCount int) {
    for _, code := range diagnostics {
        bitCount += ((int(code) >> pos) & 1)
    }

    return
}

func powerConsumptionFromDiagnostics(diagnostics []int) (gamma int, epsilon int) {
    bitCounts := bitCountsFromDiagnostics(diagnostics)
    commonCutoff := (len(diagnostics) / 2)

    for pos, count := range bitCounts {
        // count == commonCutoff is undefined behavior in original problem
        if count > commonCutoff {
            gamma |= (1 << pos)
        } else if count < commonCutoff {
            epsilon |= (1 << pos)
        }
    }

    return
}

func lifeSupportFromDiagnostics(diagnostics []int) (oxygen int, co2 int) {
    oredDiagnostics := 0
    for _, code := range diagnostics {
        oredDiagnostics |= code
    }

    startingBit := int(math.Log2(float64(oredDiagnostics)))

    var ratings []int
    ratings = ratingsFromDiagnostics(diagnostics, startingBit, oxygenGenerator)
    if len(ratings) != 1 {
        // TODO: error, and don't do check below
    }
    if len(ratings) == 1 {
        oxygen = ratings[0]
    }
    ratings = ratingsFromDiagnostics(diagnostics, startingBit, co2Scrubber)
    if len(ratings) != 1 {
        // TODO: error, and don't do check below
    }
    if len(ratings) == 1 {
        co2 = ratings[0]
    }

    return
}

type rating int

const (
    undefined rating = iota
    oxygenGenerator
    co2Scrubber
)

func ratingsFromDiagnostics(diagnostics []int, pos int, rating rating) (ratings []int) {
    if pos < 0 {
        return
    }

    bitCount := positionBitCountFromDiagnostics(diagnostics, pos)
    commonCutoff := advent.DivCeilInt(len(diagnostics), 2)
    commonCutoff = advent.MaxInt(1, commonCutoff)

    for _, code := range diagnostics {
        desiredBit := 0
        if (bitCount >= commonCutoff) {
            desiredBit = 1
        }

        // TODO: undefined rating
        if (rating == co2Scrubber) {
            desiredBit = ^desiredBit & 1
        }

        if ((code >> pos) & 1) == desiredBit {
            ratings = append(ratings, code)
        }
    }

    if len(ratings) > 1 {
        return ratingsFromDiagnostics(ratings, pos - 1, rating)
    }

    return
}
