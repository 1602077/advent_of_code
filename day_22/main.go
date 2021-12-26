package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

type Point struct {
    x, y, z int
}

type RebootStep struct {
    state bool
    start Point
    end Point
}

func (rs1 RebootStep) intersect(rs2 RebootStep, state bool) RebootStep {
    return RebootStep{
        state,
        Point{max(rs1.start.x, rs2.start.x), max(rs1.start.y, rs2.start.y), max(rs1.start.z, rs2.start.z)},
        Point{min(rs1.end.x, rs2.end.x), min(rs1.end.y, rs2.end.y), min(rs1.end.z, rs2.end.z)},
    }
}

func (rs RebootStep) isValid() bool {
    return rs.start.x <= rs.end.x && rs.start.y <= rs.end.y && rs.start.z <= rs.end.z
}

func (rs RebootStep) vol() int {
    return (abs(rs.end.x-rs.start.x) + 1) * (abs(rs.end.y-rs.start.y) + 1) * (abs(rs.end.z - rs.start.z) + 1)
}

func main() {
    part1("./input.txt")
    part2("./input.txt")
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

func part2(filename string) int {
    RebootSteps := readInput(filename, false)
    list := []RebootStep{}

    for _, rs1 := range RebootSteps {
        add := []RebootStep{}
        if rs1.state {
            add = append(add, rs1)
        }
        for _, rs2 := range list {
            if intersection := rs1.intersect(rs2, !rs2.state); intersection.isValid() {
                add = append(add, intersection)
            }
        }
        list = append(list, add...)
    }

    count := 0
    for _, rs := range list {
        if rs.state {
            count += rs.vol()
        } else {
            count -= rs.vol()
        }
    }
    fmt.Printf("PART 2 Number of cubes that are on: %v.", count)
    return count
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

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}
