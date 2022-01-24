package main

import (
    "testing"
)

func TestDomainProblem(t *testing.T) {
    caveGraph := newGraph(
        [][2]string{
            {"start", "A"},
            {"start", "b"},
            {"A", "c"},
            {"A", "b"},
            {"b", "d"},
            {"A", "end"},
            {"b", "end"},
        },
    )
    actual := pathCountWithSmallCavesOnce(caveGraph)

    if actual != 10 {
        t.Errorf("simple example: expected 10 actual %d", actual)
    }

    caveGraph = newGraph(
        [][2]string{
            {"dc", "end"},
            {"HN", "start"},
            {"start", "kj"},
            {"dc", "start"},
            {"dc", "HN"},
            {"LN", "dc"},
            {"HN", "end"},
            {"kj", "sa"},
            {"kj", "HN"},
            {"kj", "dc"},
        },
    )
    actual = pathCountWithSmallCavesOnce(caveGraph)

    if actual != 19 {
        t.Errorf("larger example: expected 19 actual %d", actual)
    }

    caveGraph = newGraph(
        [][2]string{
            {"fs", "end"},
            {"he", "DX"},
            {"fs", "he"},
            {"start", "DX"},
            {"pj", "DX"},
            {"end", "zg"},
            {"zg", "sl"},
            {"zg", "pj"},
            {"pj", "he"},
            {"RW", "he"},
            {"fs", "DX"},
            {"pj", "RW"},
            {"zg", "RW"},
            {"start", "pj"},
            {"he", "WI"},
            {"zg", "he"},
            {"pj", "fs"},
            {"start", "RW"},
        },
    )
    actual = pathCountWithSmallCavesOnce(caveGraph)

    if actual != 226 {
        t.Errorf("even larger example: expected 226 actual %d", actual)
    }
}
