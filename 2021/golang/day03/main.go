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
    fmt.Println("solution:", gamma * epsilon)
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

func powerConsumptionFromDiagnostics(diagnostics []int) (gamma int, epsilon int) {
    bitCounts := make(map[int]int)
    commonCutoff := (len(diagnostics) / 2)

    for _, code := range diagnostics {
        msb := int(math.Log2(float64(code)))

        for i := 0; i <= msb; i++ {
            bitCounts[i] += ((int(code) >> i) & 1)
        }
    }

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
