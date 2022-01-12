package main

import (
    "testing"
)

func TestDomainProblem(t *testing.T) {
    navSubsystems := []string{
        "[({(<(())[]>[[{[]{<()<>>",
        "[(()[<>])]({[<{<<[]>>(",
        "{([(<{}[<>[]}>{[]{[(<()>",
        "(((({<>}<{<{<>}{[]{[]{}",
        "[[<[([]))<([[{}[[()]]]",
        "[{[{({}]{}}([{[{{{}}([]",
        "{<[[]]>}<{[{[{[]{()[[[]",
        "[<(<(<(<{}))><([]([]()",
        "<{([([[(<>()){}]>(<<{{",
        "<{([{{}}[<[[[<>{}]]]>[]]",
    }
    actual := summedSyntaxErrorScores(navSubsystems)

    if actual != 26397 {
        t.Errorf("example: expected 26397 actual %d", actual)
    }

    actual = medianAutocompleteScore(navSubsystems)

    if actual != 288957 {
        t.Errorf("example: expected 288957 actual %d", actual)
    }
}
