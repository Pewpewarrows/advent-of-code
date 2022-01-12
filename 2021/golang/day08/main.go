package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    var data []displayPattern
    advent.Execute(scanInputData, &data)

    count := uniqueSegmentedDigitCount(data)
    fmt.Println("part one:", count)

    sum := summingDisplayOutput(data)
    fmt.Println("part two:", sum)
}

type displayPattern struct {
    signalPatterns [10]string
    displayOutput [4]string
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    data := *inputDataPtr.(*[]displayPattern)

    for scanner.Scan() {
        patterns := [10]string{}
        output := [4]string{}
        halves := strings.Split(scanner.Text(), "|")

        for i, text := range strings.Fields(halves[0]) {
            patterns[i] = text
        }

        for i, text := range strings.Fields(halves[1]) {
            output[i] = text
        }

        data = append(data, displayPattern{
            patterns,
            output,
        })
    }

    *inputDataPtr.(*[]displayPattern) = data
}

func sevenSegmentSignalsByNumber() map[int]int {
    return map[int]int{
        0: 6,
        1: 2,
        2: 5,
        3: 5,
        4: 4,
        5: 5,
        6: 6,
        7: 3,
        8: 7,
        9: 6,
    }
}

func uniqueSegmentedDigitCount(data []displayPattern) (count int) {
    // TODO: compute this dynamically
    uniqueSegmentedDigits := []int{
        sevenSegmentSignalsByNumber()[1],
        sevenSegmentSignalsByNumber()[4],
        sevenSegmentSignalsByNumber()[7],
        sevenSegmentSignalsByNumber()[8],
    }

    for _, pattern := range data {
        for _, digit := range pattern.displayOutput {
            for _, i := range uniqueSegmentedDigits {
                if len(digit) == i {
                    count++
                    break
                }
            }
        }
    }

    return
}

func summingDisplayOutput(data []displayPattern) (sum int) {
    for _, pattern := range data {
        signalPatternsBySegments := make(map[int][]string)
        // TODO: these two should be a bimap
        segmentForSignal := make(map[rune]string)
        signalForSegment := make(map[string]rune)
        // TODO: these two should be a bimap
        numberForSignalPattern := make(map[string]int)
        signalPatternForNumber := make(map[int]string)

        for _, s := range pattern.signalPatterns {
            signalPatternsBySegments[len(s)] = append(signalPatternsBySegments[len(s)], s)
        }

        // known uniques
        for _, x := range []int{1, 4, 7, 8} {
            s := advent.SortedString(signalPatternsBySegments[sevenSegmentSignalsByNumber()[x]][0])
            numberForSignalPattern[s] = x
            signalPatternForNumber[x] = s
        }

        // top: 7(3) - 1(2)
        runes := runesFromStringDiff(signalPatternForNumber[7], signalPatternForNumber[1])
        if len(runes) != 1 {
            // TODO: handle this
        }
        segmentForSignal[runes[0]] = "top"
        signalForSegment["top"] = runes[0]

        // bottom: if a (6) one has all in common with 4(4), it must be 9(6), exclude top to get this signal
        for _, s := range signalPatternsBySegments[6] {
            runes = runesFromStringDiff(s, signalPatternForNumber[4])
            if len(runes) != (6 - len(signalPatternForNumber[4])) {
                continue
            }

            s = advent.SortedString(s)
            numberForSignalPattern[s] = 9
            signalPatternForNumber[9] = s
            runes = runesFromStringDiff(string(runes), string(signalForSegment["top"]))
            if len(runes) != 1 {
                // TODO: handle this
            }
            segmentForSignal[runes[0]] = "bottom"
            signalForSegment["bottom"] = runes[0]
            break
        }

        // bottom-left: if a (6) one has exactly two in common with 7(3), it must be 6(6), and this is the exclusive signal it has with 9(6)
        for _, s := range signalPatternsBySegments[6] {
            runes = runesFromStringDiff(s, signalPatternForNumber[7])
            if len(runes) != (6 - 2) {
                continue
            }

            s = advent.SortedString(s)
            numberForSignalPattern[s] = 6
            signalPatternForNumber[6] = s
            runes = runesFromStringDiff(s, signalPatternForNumber[9])
            if len(runes) != 1 {
                // TODO: handle this
            }
            segmentForSignal[runes[0]] = "bottom-left"
            signalForSegment["bottom-left"] = runes[0]
            break
        }

        // can now deduce which is 0(6)
        for _, s := range signalPatternsBySegments[6] {
            s = advent.SortedString(s)
            if (s == signalPatternForNumber[6]) || (s == signalPatternForNumber[9]) {
                continue
            }

            numberForSignalPattern[s] = 0
            signalPatternForNumber[0] = s
            break
        }

        // if a (5) one has exactly two in common with 4(4), it must be 2(5)
        for _, s := range signalPatternsBySegments[5] {
            runes = runesFromStringDiff(s, signalPatternForNumber[4])
            if len(runes) != (5 - 2) {
                continue
            }

            s = advent.SortedString(s)
            numberForSignalPattern[s] = 2
            signalPatternForNumber[2] = s
            break
        }

        // if 2(5) has exactly two in common with a (5) one, it must be 5(5)
        for _, s := range signalPatternsBySegments[5] {
            runes = runesFromStringDiff(signalPatternForNumber[2], s)
            if len(runes) != 2 {
                continue
            }

            s = advent.SortedString(s)
            numberForSignalPattern[s] = 5
            signalPatternForNumber[5] = s
            break
        }

        // can now deduce which is 3(5)
        for _, s := range signalPatternsBySegments[5] {
            s = advent.SortedString(s)
            if (s == signalPatternForNumber[2]) || (s == signalPatternForNumber[5]) {
                continue
            }

            numberForSignalPattern[s] = 3
            signalPatternForNumber[3] = s
            break
        }

        // TODO: finish filling these out if ever necessary:
        // top-right: see bottom-left, this is the exclusive signal 9(6) has with 6(6)
        // bottom-right: remaining signal in 7(3) that's not top or top-right
        // mid: 6(6) - 0(6)
        // top-left: remaining unsolved signal

        digits := ""
        for _, digit := range pattern.displayOutput {
            if n, ok := numberForSignalPattern[advent.SortedString(digit)]; ok {
                digits += fmt.Sprint(n)
            }
        }

        i, err := strconv.ParseInt(digits, 10, 32)
        if err != nil {
            // TODO: handle me
        }
        // TODO: why is i here an int64?
        sum += int(i)
    }

    return
}

func runesFromStringDiff(a, b string) (diff []rune) {
    for _, c := range a {
        if !strings.Contains(b, string(c)) {
            diff = append(diff, c)
        }
    }

    return
}
