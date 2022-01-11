package main

import (
	"testing"
)

func TestCheapestFuelForAlignment(t *testing.T) {
	crabYs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	actual := cheapestFuelForAlignment(crabYs)

	if actual != 37 {
		t.Errorf("example: expected 37 actual %d", actual)
	}

	actual = cheapestIncrementalFuelForAlignment(crabYs)

	if actual != 168 {
		t.Errorf("example incremental: expected 168 actual %d", actual)
	}
}
