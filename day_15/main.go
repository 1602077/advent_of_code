package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "sort"
    "strings"
)

type point struct {
    x int
    y int
}

func (p point) add(p2 point) point {
    return point{p.x + p2.x, p.y + p2.y}
}


type gridPoint struct {
    pos point
    value int
}

type grid struct {
    point map[point]gridPoint
    x int
    y int
}

type state struct {
    score int
    pos point
}

var dirs = []point{
    {x: 0, y: 1},
    {x: 0, y: -1},
    {x: 1, y: 0},
    {x: -1, y: 0},
}

func main() {
    // part1("./input_test.txt")
    part1("./input.txt")
}

func part1(filename string) int {
    grid := readInput(filename)

    best := make(map[point]int)
    queue := make([]state, 0)

    start := point{x: 0, y: 0}
    best[start] = 0
    queue = append(queue, state{0, start})

    for len(queue) > 0 {
        currState := queue[0]
        queue = queue[1:]
        next := []gridPoint{}

        for _, d := range dirs {
            n, ok := grid.point[currState.pos.add(d)]
            if ok {
                if b, ok := best[n.pos]; ok {
                    vnext := currState.score + n.value
                    if b > vnext {
                        best[n.pos] = vnext
                        next = append(next, n)
                    }

                } else {
                    best[n.pos] = currState.score + n.value
                    next = append(next, n)
                }
            }
        }
        sort.Slice(next, func(i, j int) bool {return next[i].value < next[j].value})

        for _, gp := range next {
            queue = append(queue, state{currState.score + gp.value, gp.pos})
        }
    }
    best_score := best[point{x: int(grid.x - 1) , y: int(grid.y - 2)}]
    fmt.Printf("PART 1 Best score to end of grid: %v.\n", best_score)
    return best_score
}

func readInput(filename string) grid {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    input := strings.Split(string(data), "\n")
    g := grid{make(map[point]gridPoint), len([]rune(input[0])), len(input)}

    for i, ln := range input {
        for j, num := range ln {
            g.point[point{x: j, y: i}] = gridPoint{point{x: j, y: i}, int(num - '0')}

        }
    }
    return g
}
