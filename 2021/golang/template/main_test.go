package main

import (
    "testing"
)

func TestDomainProblem(t *testing.T) {
    course := []int{7}
    actual := domainProblem(course)

    if actual != 42 {
        t.Errorf("simple: expected 42 actual %d", actual)
    }
}
