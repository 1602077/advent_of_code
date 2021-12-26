package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

type RebootStep struct {
    state bool
    start Point
    end Point
}

type Point struct {
    x, y, z int
}

func main() {
    part1("./input.txt")
}

func part1(filename string) int {
    RebootSteps := readInput(filename, true)

    cubes := map[Point]bool{}

    for _, rs := range RebootSteps {
        for i := rs.start.x; i <= rs.end.x; i++ {
            for j := rs.start.y; j <= rs.end.y; j++ {
                for k := rs.start.z; k <= rs.end.z; k++ {
                    cubes[Point{i,j,k}] = rs.state
                }
            }
        }
    }
    numCubes := 0
    for _, state := range cubes {
        if state {
            numCubes++
        }
    }
    fmt.Printf("PART 1 Number of cubes on: %v.\n", numCubes)
    return numCubes
}

func readInput(filename string, limitRange bool) []RebootStep {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    var RebootSteps []RebootStep
    var boolean string
    var x1, x2, y1, y2, z1, z2 int

    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &boolean, &x1, &x2, &y1, &y2, &z1, &z2)
        state := false
        if boolean == "on" {
            state = true
        }
        if limitRange {
            rs := RebootStep{
                state,
                Point{max(x1, -50), max(-50, y1), max(-50, z1)},
                Point{min(x2, 50), min(50, y2), min(50, z2)},
            }
            RebootSteps = append(RebootSteps, rs)
        } else {
            rs := RebootStep{
                state,
                Point{x1, y1, z1},
                Point{x2, y2, z2},
            }
            RebootSteps = append(RebootSteps, rs)
        }
    }
    return RebootSteps
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

