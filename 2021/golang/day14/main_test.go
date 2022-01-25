package main

import (
    "testing"
)

func TestElementCountSpread(t *testing.T) {
    template := "NNCB"
    rules := map[string]rune{
        "CH": 'B',
        "HH": 'N',
        "CB": 'H',
        "NH": 'C',
        "HB": 'C',
        "HC": 'B',
        "HN": 'C',
        "NN": 'C',
        "BH": 'H',
        "NC": 'B',
        "NB": 'B',
        "BN": 'B',
        "BB": 'N',
        "BC": 'B',
        "CC": 'N',
        "CN": 'C',
    }
    actual := elementCountSpread(template, rules, 10)

    if actual != 1588 {
        t.Errorf("simple: expected 1588 actual %d", actual)
    }
}
