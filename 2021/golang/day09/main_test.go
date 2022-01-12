package main

import (
	"testing"
)

func TestDomainProblem(t *testing.T) {
	heightMap := [][]int{
		[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	actual := summedLowPointRiskLevels(heightMap)

	if actual != 15 {
		t.Errorf("example sum: expected 15 actual %d", actual)
	}

	actual = productOfLargestBasins(heightMap, 3)

	if actual != 1134 {
		t.Errorf("example product: expected 1134 actual %d", actual)
	}
}
