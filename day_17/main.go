package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    part1("./input.txt")
}

func part1(filename string) int {
    _, y := readInput(filename)
    var minY int
    for _, v := range y {
        if minY > v {
            minY = v
        }
    }
    maxY  := (minY + 1) * (minY) / 2
    fmt.Printf("PART 1 Highest y position while still on trajectory: %v.\n", maxY)
    return maxY
}

func part2(filename string) {
    _, y := readInput(filename)
    var minY int
    for _, v := range y {
        if minY > v {
            minY = v
        }
    }
    maxY  := (minY + 1) * (minY) / 2
    fmt.Print(maxY)
}

func readInput(filename string) ([2]int, [2]int) {
    input, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    var x, y [2]int
    fmt.Sscanf(string(input), "target area: x=%d..%d, y=%d..%d", &x[0], &x[1], &y[0], &y[1])
    return x, y
}
