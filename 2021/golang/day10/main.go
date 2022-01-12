package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "strings"
)

func main() {
    var navSubsystems []string
    advent.Execute(scanInputData, &navSubsystems)

    sum := summedSyntaxErrorScores(navSubsystems)
    fmt.Println("part one:", sum)

    median := medianAutocompleteScore(navSubsystems)
    fmt.Println("part two:", median)
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    navSubsystems := *inputDataPtr.(*[]string)

    for scanner.Scan() {
        navSubsystems = append(navSubsystems, scanner.Text())
    }

    *inputDataPtr.(*[]string) = navSubsystems
}

func summedSyntaxErrorScores(navSubsystems []string) (sum int) {
    pointsByIllegalRune := map[rune]int {
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137,
    }

    illegalRunes, _ := navSyntaxErrors(navSubsystems)

    for _, r := range illegalRunes {
        sum += pointsByIllegalRune[r]
    }

    return
}

func navSyntaxErrors(navSubsystems []string) (illegalRunes []rune, autocompletes []string) {
    // TODO: these two should be in a bimap
    chunkCloserByOpener := map[rune]rune {
        '(': ')',
        '[': ']',
        '{': '}',
        '<': '>',
    }

    chunkOpenerByCloser := map[rune]rune {
        ')': '(',
        ']': '[',
        '}': '{',
        '>': '<',
    }

subsystemLoop:
    for _, nav := range navSubsystems {
        var chunkStack []rune
        var sb strings.Builder

        for _, r := range nav {
            // TODO: handle illegal opener
            if (r == '(') || (r == '[') || (r == '{') || (r == '<') {
                chunkStack = append(chunkStack, r)
                continue
            }

            opener, ok := chunkOpenerByCloser[r]

            if !ok {
                // TODO: handle illegal closer
            }

            // pop
            if len(chunkStack) == 0 {
                illegalRunes = append(illegalRunes, r)
                continue subsystemLoop
            }

            var top rune
            top, chunkStack = chunkStack[len(chunkStack) - 1], chunkStack[:len(chunkStack) - 1]

            if top != opener {
                illegalRunes = append(illegalRunes, r)
                continue subsystemLoop
            }
        }

        for {
            // pop
            // TODO: refactor into pop function
            if len(chunkStack) == 0 {
                break
            }

            var top rune
            top, chunkStack = chunkStack[len(chunkStack) - 1], chunkStack[:len(chunkStack) - 1]

            sb.WriteRune(chunkCloserByOpener[top])
        }

        if sb.Len() != 0 {
            autocompletes = append(autocompletes, sb.String())
        }
    }

    return
}

func medianAutocompleteScore(navSubsystems []string) int {
    pointsByChunkCloser := map[rune]int {
        ')': 1,
        ']': 2,
        '}': 3,
        '>': 4,
    }

    _, autocompletes := navSyntaxErrors(navSubsystems)
    var scores []int

    for _, text := range autocompletes {
        score := 0
        for _, r := range text {
            score *= 5
            score += pointsByChunkCloser[r]
        }

        scores = append(scores, score)
    }

    return int(advent.MedianInts(scores))
}
