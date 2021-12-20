package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    part1("./input.txt")
    part2("./input.txt")
}

func part1(filename string) int {
    _, y := readInput(filename)
    maxY  := (y[0] + 1) * (y[0]) / 2
    fmt.Printf("PART 1 Highest y position while still on trajectory: %v.\n", maxY)
    return maxY
}

func part2(filename string) int {
    x, y := readInput(filename)
    vXmin, vXmax := 1, x[1]
    vYmin, vYmax := y[0], (y[0] + 1) * (y[0]) / 2

    hitCount := 0
    for vX := vXmin; vX <= vXmax; vX++ {
        for vY := vYmin; vY <= vYmax; vY++ {
            hit := false
            vel := [2]int{vX, vY}
            pos := [2]int{0, 0}
            for {
                if x[0] <= pos[0] && pos[0] <= x[1] && y[0] <= pos[1] && pos[1] <= y[1] {
                    hit = true
                    break
                }
                if pos[0] > x[1] || pos[1] < y[0] {
                    break
                }
                pos = [2]int{pos[0] + vel[0], pos[1] + vel[1]}

                if vel[0] > 0 {
                    vel[0]--
                    vel[1]--
                } else if vel[0] < 0 {
                    vel[0]++
                    vel[1]--
                } else {
                    vel[1]--
                }
            }
            if hit {
                hitCount++
            }

        }
    }
    fmt.Printf("PART 2 Hit Count: %v.\n", hitCount)
    return hitCount
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
