package main

import (
    "strings"
    "strconv"
    "testing"
)

var testInput = `607
618
618
617
647
716
769
792`

var testAnswer = 5

func TestP2Solution(t *testing.T) {
    type test struct {
        input string
        answer int
    }
    test1 := test{input: testInput, answer: 5}

    // data := readInput(test1.input)
    splitContent := strings.Split(test1.input, "\n")

    depths := make([]int, len(splitContent))

    for i, splitContent := range splitContent {
        depths[i], _ = strconv.Atoi(splitContent)
    }

    result := partTwoSolver(depths)

    if result != test1.answer {
        t.Errorf("ERROR: Expected %d, got %d", test1.answer, result)
    }

}
