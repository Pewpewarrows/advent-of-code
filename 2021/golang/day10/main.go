package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
)

func main() {
    var navSubsystems []string
    advent.Execute(scanInputData, &navSubsystems)

    sum := summedSyntaxErrorScores(navSubsystems)

    fmt.Println("solution:", sum)
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

    chunkOpenerByCloser := map[rune]rune {
        ')': '(',
        ']': '[',
        '}': '{',
        '>': '<',
    }

    for _, nav := range navSubsystems {
        var chunkStack []rune

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
                sum += pointsByIllegalRune[r]
                continue
            }

            var top rune
            top, chunkStack = chunkStack[len(chunkStack) - 1], chunkStack[:len(chunkStack) - 1]

            if top != opener {
                sum += pointsByIllegalRune[r]
            }
        }
    }

    return
}
