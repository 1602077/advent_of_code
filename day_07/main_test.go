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
    }
    for _, test := range tests {
        r := part1(test.in)
        if r != test.ans {
            t.Errorf("ERROR: Expected %d, got %d.\n", test.ans, r)
        }
    }
}
