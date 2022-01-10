package main

import (
	"testing"
)

func TestVentOverlapCount(t *testing.T) {
	ventLines := []ventLine{
		{0, 9, 5, 9},
		{8, 0, 0, 8},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
		{7, 0, 7, 4},
		{6, 4, 2, 0},
		{0, 9, 2, 9},
		{3, 4, 1, 4},
		{0, 0, 8, 8},
		{5, 5, 8, 2},
	}
	actual := ventOverlapCount(ventLines)

	if actual != 12 {
		t.Errorf("example: expected 12 actual %d", actual)
	}
}
