package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "fmt"
    "math"
    "strconv"
    "strings"
)

func main() {
    var game bingoGame
    advent.Execute(scanInputData, &game)

    score := winningBingoScore(game)
    fmt.Println("part one:", score)

    score = lastWinningBingoScore(game)
    fmt.Println("part two:", score)
}

type bingoGame struct {
    balls []int
    boards []bingoBoard
}

const boardSize = 5
// TODO: const totalBoardSpaces = boardSize * boardSize
var allWinningBitmasks = winningBitmasks(boardSize)

func winningBitmasks(boardSize int) (bitmasks []int) {
    // TODO: this only works up to a certain size based on int type
    rowMask := int(math.Pow(2, float64(boardSize)) - 1)

    for i := 0; i < (boardSize * boardSize); i += boardSize {
        bitmasks = append(bitmasks, rowMask << i)
    }

    var verticalMask int

    for i := 0; i < boardSize; i++ {
        verticalMask = 0
        for j := 0; j < boardSize; j++ {
            verticalMask |= ((1 << i) << (j * boardSize))
        }
        bitmasks = append(bitmasks, verticalMask)
    }

    // TODO: only if debug
    // for _, bitmask := range bitmasks {
    //     fmt.Printf("%025b\n", bitmask)
    // }

    return
}

type bingoBoard struct {
    numbers [boardSize * boardSize]int
    markedBitmask int
}

func (b *bingoBoard) String() string {
    // TODO: the look of the bitmask is inverted from the actual list of
    // numbers due to indexes
    return fmt.Sprintf("%025b", b.markedBitmask)
}

func (b *bingoBoard) markBall(ball int) {
    for i, n := range b.numbers {
        if n == ball {
            b.markedBitmask |= (1 << i)
            break
        }
    }
}

func (b *bingoBoard) isWinning() (winning bool) {
    for _, bitmask := range allWinningBitmasks {
        if (b.markedBitmask & bitmask) == bitmask {
            return true
        }
    }

    return false
}

func (b *bingoBoard) score(winningBall int) (score int) {
    for i, n := range b.numbers {
        if (b.markedBitmask & (1 << i)) == 0 {
            score += n
        }
    }

    score *= winningBall

    return
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    game := *inputDataPtr.(*bingoGame)
    state := initial
    var numbers [boardSize * boardSize]int
    var rowNum int

    for scanner.Scan() {
        switch state {
        case initial:
            // could instead do something like:
            // bufio.NewScanner(strings.NewReader(scanner.Text()))
            // with a custom function to pass to its .Split() and then do a new
            // .Scan() but that's a bit heavy-handed
            for _, text := range strings.Split(scanner.Text(), ",") {
                i, err := strconv.ParseInt(text, 10, 32)
                if err != nil {
                    // TODO: handle me
                }
                // TODO: why is i here an int64?
                game.balls = append(game.balls, int(i))
            }
            state = empty
        case empty:
            // skip the newline, don't consume scanner.Text()
            state = boardRow
        case boardRow:
            // can do the same strings.Split() as above, but here's another
            // quick way to do the same thing
            for j, text := range strings.Fields(scanner.Text()) {
                i, err := strconv.ParseInt(text, 10, 32)
                if err != nil {
                    // TODO: handle me
                }
                // TODO: why is i here an int64?
                numbers[j + (rowNum * boardSize)] = int(i)
            }

            rowNum++

            if rowNum == boardSize {
                game.boards = append(game.boards, bingoBoard{
                    numbers,
                    0,
                })

                rowNum = 0
                numbers = [boardSize * boardSize]int{}
                state = empty
            }
        default:
            // TODO: handle me
        }
    }

    *inputDataPtr.(*bingoGame) = game
}

type scanState int

const (
    undefined scanState = iota
    initial
    empty
    boardRow
)

func winningBingoScore(game bingoGame) (score int) {
    var winner *bingoBoard
    var winningBall int

    for _, ball := range game.balls {
        for i := range game.boards {
            game.boards[i].markBall(ball)

            if game.boards[i].isWinning() {
                winner = &game.boards[i]
                break
            }
        }

        if (winner != nil) {
            winningBall = ball
            break
        }
    }

    if (winner == nil) {
        // TODO: error
    }

    score = winner.score(winningBall)

    return
}

func lastWinningBingoScore(game bingoGame) (score int) {
    var winner *bingoBoard
    var winningBall int
    boardCount := len(game.boards)
    var winningBoardCount int

    for _, ball := range game.balls {
        winningBoardCount = 0

        for i := range game.boards {
            if game.boards[i].isWinning() {
                winningBoardCount++
                continue
            }

            game.boards[i].markBall(ball)

            if game.boards[i].isWinning() {
                winningBoardCount++
                winner = &game.boards[i]
            }
        }

        if winningBoardCount == boardCount {
            winningBall = ball
            break
        }
    }

    if (winner == nil) {
        // TODO: error
    }

    score = winner.score(winningBall)

    return
}
