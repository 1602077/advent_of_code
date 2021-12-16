package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main() {
    part1("./input.txt")
    part2("./input.txt")
}

var (
    dirs = []point{
        {x: 0, y: -1},
        {x: -1, y: -1},
        {x: -1, y: 0},
        {x: 1, y: 1},
        {x: 0, y: 1},
        {x: -1, y: 1},
        {x: 1, y: 0},
        {x: 1, y: -1},
    }
)

var flashCount int

type point struct {
    x,y int
}

type oct struct {
    energy int
    flash bool
}

func part1(filename string) int {
    grid := readInput(filename)
    for k := 0; k < 100; k++ {
        for y := 0; y < len(grid); y++ {
            for x:= 0; x < len(grid[y]); x++ {
                if !grid[y][x].flash {
                    grid[y][x].energy++
                    if grid[y][x].energy == 10 {
                        flashPoint(point{x: x, y: y}, k, grid)
                    }
                }
            }
        }
        for y := 0; y < len(grid); y++ {
            for x := 0; x < len(grid[y]); x++ {
                if grid[y][x].flash {
                    grid[y][x].flash = false
                }
            }
        }
    }
    fmt.Printf("Part 1 Total number of flashes: %v.\n", flashCount)
    return flashCount
}

func part2(filename string) int {
    grid := readInput(filename)
    sync_step := 0
    // keep iterating until all flash
    for k := 0; ; k++ {
        for y := 0; y < len(grid); y++ {
            for x:= 0; x < len(grid[y]); x++ {
                if !grid[y][x].flash {
                    grid[y][x].energy++
                    if grid[y][x].energy == 10 {
                        flashPoint(point{x: x, y: y}, k, grid)
                    }
                }
            }
        }
        allFlash := true
        for y := 0; y < len(grid); y++ {
            for x := 0; x < len(grid[y]); x++ {
                if grid[y][x].flash {
                    grid[y][x].flash = false
                }
                if grid[y][x].energy != 0 {
                    allFlash = false
                }
            }
        }
        if allFlash {
            sync_step = k +1
            break
        }
    }
    fmt.Printf("Part 2 Synchronisation first achieved at step: %v.\n", sync_step)
    return sync_step
}
func flashPoint(currPoint point, currStep int, grid [10][10]*oct) {
    flashCount++
    grid[currPoint.y][currPoint.x].energy = 0
    grid[currPoint.y][currPoint.x].flash = true
    for _, d := range dirs {
        p := point{x: currPoint.x + d.x, y: currPoint.y + d.y}
        if p.x >= 0 && p.x < len(grid[currPoint.y]) && p.y >= 0 && p.y < len(grid) {
            if !grid[p.y][p.x].flash {
                grid[p.y][p.x].energy++
                if grid[p.y][p.x].energy == 10 {
                    flashPoint(p, currStep, grid)
                }
            }
        }
    }
}

func readInput(filename string) [10][10]*oct {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    inStr := strings.Split(string(data), "\n")

    grid := [10][10]*oct{}
    for i, val := range inStr {
        for j, num := range val {
            n := num - '0'
            grid[i][j] = &oct{energy: int(n)}
        }
    }
    return grid
}
