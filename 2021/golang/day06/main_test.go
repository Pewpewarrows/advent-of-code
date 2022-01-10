package main

import (
    "testing"
)

func TestSimulatingLanternfishPopulation(t *testing.T) {
    fishSchool := []int{3, 4, 3, 1, 2}
    actual := simulatingLanternfishPopulation(fishSchool, 80)

    if actual != 5934  {
        t.Errorf("simple: expected 5934 actual %d", actual)
    }
}
