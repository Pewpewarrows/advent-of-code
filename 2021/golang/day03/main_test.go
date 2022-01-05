package main

import (
    "testing"
)

func TestDomainProblem(t *testing.T) {
    diagnostics := []int{0b11111}
    gamma, epsilon := powerConsumptionFromDiagnostics(diagnostics)

    if (gamma != 31) || (epsilon != 0) {
        t.Errorf("simple: expected (31, 0) actual (%d, %d)", gamma, epsilon)
    }

    diagnostics = []int{
        0b00100,
        0b11110,
        0b10110,
        0b10111,
        0b10101,
        0b01111,
        0b00111,
        0b11100,
        0b10000,
        0b11001,
        0b00010,
        0b01010,
    }
    gamma, epsilon = powerConsumptionFromDiagnostics(diagnostics)

    if (gamma != 22) || (epsilon != 9) {
        t.Errorf("example: expected (22, 9) actual (%d, %d)", gamma, epsilon)
    }

    diagnostics = []int{}
    gamma, epsilon = powerConsumptionFromDiagnostics(diagnostics)

    if (gamma != 0) || (epsilon != 0) {
        t.Errorf("empty: expected (0, 0) actual (%d, %d)", gamma, epsilon)
    }
}
