package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "strings"
)

func main() {
    var p puzzle
    advent.Execute(scanInputData, &p)

    solution := elementCountSpread(p.template, p.rules, 10)
    fmt.Println("part one:", solution)

    solution = elementCountSpread(p.template, p.rules, 40)
    fmt.Println("part two:", solution)
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
    polymer := make(map[string]int)

    for i, r := range template {
        if i == (len(template) - 1) {
            break
        }

        polymer[string([]rune{r, rune(template[i + 1])})]++
    }

    for i := 0; i < stepCount; i++ {
        nextPolymer := make(map[string]int)

        for pair, c := range polymer {
            nextPolymer[string([]rune{rune(pair[0]), rules[pair]})] += c
            nextPolymer[string([]rune{rules[pair], rune(pair[1])})] += c
        }

        polymer = nextPolymer
    }

    elementCounts := make(map[rune]int)
    for pair, c := range polymer {
        elementCounts[rune(pair[0])] += c
        elementCounts[rune(pair[1])] += c
    }

    elementCounts[rune(template[0])]++
    elementCounts[rune(template[len(template) - 1])]++

    for i := range elementCounts {
        elementCounts[i] /= 2
    }

    largestCount := elementCounts[rune(template[0])]
    smallestCount := elementCounts[rune(template[0])]

    for _, c := range elementCounts {
        if c > largestCount {
            largestCount = c
        } else if c < smallestCount {
            smallestCount = c
        }
    }

    return largestCount - smallestCount
}
