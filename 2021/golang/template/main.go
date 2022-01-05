package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
)

func main() {
    var data []int
    advent.Execute(scanInputData, &data)

    solution := domainProblem(data)

    fmt.Println("solution:", solution)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    data := *inputDataPtr.(*[]int)
    var i int

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d", &i)
        data = append(data, i)
    }

    *inputDataPtr.(*[]int) = data
}

func domainProblem(data []int) (solution int) {
    solution = 42
    return
}
