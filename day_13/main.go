package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

type point struct {
    x int
    y int
}

type fold struct {
    dir rune
    val int
}

func main() {
    part1("./input.txt")
    part2("./input.txt")
}

func part1(filename string) int {
    points, folds := readInput(filename)
    grid := createGid(points)
    foldedGrid := foldGrid(grid, folds, true)

    count := 0
    for i := 0; i < len(foldedGrid); i++ {
        for j := 0; j < len(foldedGrid[i]); j++ {
            if foldedGrid[i][j] == "#" {
                count++
            }
        }
    }
    fmt.Printf("PART 1 Number of active pixels: %v.\n", count)
    return count
}

func part2(filename string)  {
    points, folds := readInput(filename)
    grid := createGid(points)
    foldedGrid := foldGrid(grid, folds, false)

    fmt.Println("PART 2 folded grid")
     for _, v := range foldedGrid {
        for _, s := range v {
            if s != "#" {
                fmt.Print(".")
            } else {
                fmt.Print("#")
            }
        }
        fmt.Println()
    }
}

func readInput(filename string) ([]point, []fold) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    contents := strings.Split(string(data), "\n\n")
    var points []point
    points_str := strings.Split(contents[0], "\n")
    for _, line := range points_str {
        var (
            x int
            y int
        )
        if _, err :=fmt.Sscanf(line, "%v,%v", &x, &y); err != nil {
            log.Fatal(err)
        }
        points = append(points, point{x: x, y: y})
    }

    var folds []fold
    fold_str := strings.Split(contents[1], "\n")
    for _, line := range fold_str[:len(fold_str)-1] {
        var (
            dir rune
            val int
        )
        if _, err := fmt.Sscanf(line, "fold along %c=%d", &dir, &val); err != nil {
            log.Fatal(err)
        }
        folds = append(folds, fold{dir: dir, val: val})
    }
    return points, folds
}

func createGid(points []point) [][]string {
    var grid [][]string
    maxX, maxY := 0, 0
    for _, p := range points {
        if p.x > maxX {
            maxX = p.x
        }
        if p.y > maxY {
            maxY = p.y
        }
    }

    grid = make([][]string, maxY+1)
    for i := 0; i < maxY+1; i++ {
        grid[i] = make([]string, maxX+1)
    }
    for _, p:= range points {
        grid[p.y][p.x] = "#"
    }
    return grid
}

func foldGrid(grid [][]string, folds []fold, firstFoldOnly bool) [][]string {
    for _, f := range folds {
        if f.dir == 'x' {
            var x int = f.val
            for y := 0; y < len(grid); y++ {
                for j := x; j < len(grid[y]); j++ {
                    if grid[y][j] == "#" {
                        grid[y][2*x - j] = "#"
                    }
                }
            }
            for i, v := range grid {
                v = v[:x]
                grid[i] = v
            }
        } else if f.dir == 'y' {
            var y int = f.val
            for i := y; i < len(grid); i++ {
                for x := 0; x < len(grid[i]); x++ {
                    if grid[i][x] == "#" {
                        grid[2*y-i][x] = "#"
                    }
                }
            }
            grid = grid[:y]
        }
        if firstFoldOnly {
            break
        }
    }
    return grid
}
