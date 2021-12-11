package main

import (
    "testing"
)

func TestSol(t *testing.T) {
    tests := []struct {
        in string
        ans int
    }{
        {"./input_test.txt", 37},
        {"./input_test.txt", 168},
    }
    // testing part1()
    r1 := part1(tests[0].in)
    if r1 != tests[0].ans {
        t.Errorf("ERROR: Expected %d, got %d.\n", tests[0].ans, r1)
    }
    // testing part2()
    r2 := part2(tests[1].in)
    if r2 != tests[1].ans {
        t.Errorf("ERROR: Expected %d, got %d.\n", tests[1].ans, r2)
    }
}
