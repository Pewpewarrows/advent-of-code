package main

import (
    "testing"
)

func TestCaveGraphPathCount(t *testing.T) {
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
    actual := caveGraph.pathCount(1, 1)
    if actual != 10 {
        t.Errorf("simple example no small revisits: expected 10 actual %d", actual)
    }
    actual = caveGraph.pathCount(1, 2)
    if actual != 36 {
        t.Errorf("simple example 1/2 small revisits: expected 36 actual %d", actual)
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
    actual = caveGraph.pathCount(1, 1)
    if actual != 19 {
        t.Errorf("larger example no small revisits: expected 19 actual %d", actual)
    }
    actual = caveGraph.pathCount(1, 2)
    if actual != 103 {
        t.Errorf("larger example 1/2 small revisits: expected 103 actual %d", actual)
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
    actual = caveGraph.pathCount(1, 1)
    if actual != 226 {
        t.Errorf("even larger example no small revisits: expected 226 actual %d", actual)
    }
    actual = caveGraph.pathCount(1, 2)
    if actual != 3509 {
        t.Errorf("even larger example 1/2 small revisits: expected 3509 actual %d", actual)
    }
}
