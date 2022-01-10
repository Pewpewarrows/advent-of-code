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

    // TODO: why did fishSchool get modified above?
    fishSchool = []int{3, 4, 3, 1, 2}
    actual = estimatingLanternfishPopulation(fishSchool, 256)

    if actual != 26984457539  {
        t.Errorf("simple: expected 26984457539 actual %d", actual)
    }
}
