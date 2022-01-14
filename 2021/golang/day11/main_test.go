package main

import (
	"testing"
)

func TestDomainProblem(t *testing.T) {
	octopusGrid := [10][10]int{
		[10]int{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		[10]int{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		[10]int{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		[10]int{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		[10]int{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		[10]int{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		[10]int{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		[10]int{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		[10]int{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		[10]int{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}
	actual := flashCountSimulation(octopusGrid, 100)

	if actual != 1656 {
		t.Errorf("example: expected 1656 actual %d", actual)
	}
}
