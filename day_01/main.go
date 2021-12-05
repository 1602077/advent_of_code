package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func main() {
    // Read in inputs
    depths := readInput("./input.txt")
    fmt.Printf("The first 10 depths are: %v.\n", depths[:10])

    // Part 1 solution
    increaseCount := partOneSolver(depths)
    fmt.Printf("Number of times a depth measurement increases point to point: %v.\n", increaseCount)

    windowedCount := partTwoSolver(depths)
    fmt.Printf("Number of times a depth measurement increases over a 3-day window: %v.\n", windowedCount)
}

func readInput(filename string) []int {
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    contents := string(file)
    splitContent := strings.Split(contents, "\n")

    depths := make([]int, len(splitContent))

    for i, splitContent := range splitContent {
        depths[i], _ = strconv.Atoi(splitContent)
    }

    return depths
}

func partOneSolver(depths []int) int {
    count := 0
    var prev int

    for i, curr:= range depths {
        if i > 0 && curr - prev > 0 {
            count++
        }
        prev=curr
    }
    return count
}

func partTwoSolver(depths []int) int {
    // Create point measurements to sliding 3 point sums
    sums := []int{}

    for i := 0; i < len(depths) - 2; i++ {
        sum := depths[i] + depths[i+1] + depths[i+2]
        sums = append(sums, sum)
    }

    count := 0
    var prev int
    for i, curr := range sums {
        if i > 0 && curr - prev > 0 {
            count++
        }
        prev=curr
    }
    return count
}
