package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "bytes"
    "fmt"
    "strings"
)

func main() {
    var p puzzle
    advent.Execute(scanInputData, &p)

    solution := elementCountSpread(p.template, p.rules, 10)

    fmt.Println("solution:", solution)
}

type puzzle struct {
    template string
    rules map[string]rune
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    p := *inputDataPtr.(*puzzle)
    rules := make(map[string]rune)

    // TODO: state enum
    isTemplateRead := false

    for scanner.Scan() {
        if scanner.Text() == "" {
            isTemplateRead = true
            continue
        }

        if isTemplateRead {
            halves := strings.Split(scanner.Text(), "->")
            // TODO: assumes right half is a single rune
            rules[strings.TrimSpace(halves[0])] = rune(strings.TrimSpace(halves[1])[0])
        } else {
            fmt.Sscanf(scanner.Text(), "%s", &p.template)
        }
    }

    p.rules = rules

    *inputDataPtr.(*puzzle) = p
}

func elementCountSpread(template string, rules map[string]rune, stepCount int) int {
    polymer := template

    for i := 0; i < stepCount; i++ {
        var b bytes.Buffer

        for j, c := range polymer {
            if j == (len(polymer) - 1) {
                b.WriteRune(c)
                break
            }

            b.WriteRune(c)
            b.WriteRune(rules[string([]rune{c, rune(polymer[j + 1])})])
        }

        polymer = b.String()
    }

    elementCounts := make(map[rune]int)
    for _, r := range polymer {
        elementCounts[r]++
    }

    largestCount := elementCounts[rune(polymer[0])]
    smallestCount := elementCounts[rune(polymer[0])]

    for _, c := range elementCounts {
        if c > largestCount {
            largestCount = c
        } else if c < smallestCount {
            smallestCount = c
        }
    }

    return largestCount - smallestCount
}
