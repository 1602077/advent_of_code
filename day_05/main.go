package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func main() {
    part1()
    part2()
}

func part1() {
    coords := parseInputFile()
    var grid [1000][1000]uint8

    for _, c := range coords {
        if c.x1 == c.x2 {
            ls :=  createHVLine(c.y1, c.y2)
            for _, p := range ls {
                grid[c.x1][p]++
            }
        } else if c.y2 == c.y1 {
            ls := createHVLine(c.x1, c.x2)
            for _, p := range ls {
                grid[p][c.y1]++
            }
        }
    }

    numInter := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] > 1 {
                numInter++
            }
        }
    }
    fmt.Printf("PART 1 Number of intersections: %v.\n", numInter)
}

func part2() {
    coords := parseInputFile()
    var grid [1000][1000]uint8

    for _, c := range coords {
        // handles horizontal & vertical lines
        if c.x1 == c.x2 {
            ls :=  createHVLine(c.y1, c.y2)
            for _, p := range ls {
                grid[c.x1][p]++
            }
        } else if c.y2 == c.y1 {
            ls := createHVLine(c.x1, c.x2)
            for _, p := range ls {
                grid[p][c.y1]++
            }
         } else if (c.x2 - c.x1) == (c.y2 - c.y1){
            // TODO: handle diagonal lines
            diagLine := createDiagLine(c)
            for _, p := range diagLine {
                grid[p[0]][p[1]]++
                fmt.Print(p[0], " ", p[1], " \n")
            }
        }
    }

    numInter := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] > 1 {
                numInter++
            }
        }
    }
    fmt.Printf("PART 2 Number of intersections: %v.\n", numInter)
}

type coord struct {
    x1 int
    y1 int
    x2 int
    y2 int
}

func parseInputFile() []coord  {
    data, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        log.Fatal(err)
    }

    var lineSegs []coord
    lines := strings.Split(string(data), "\n")
    for _, line := range lines[:len(lines)-1] {
        lineSegs = append(lineSegs, parseCoord(line))
    }
    return lineSegs
}

func parseCoord(input string) coord {
    parts := strings.Fields(input)
    // Convert str -> []str -> []int.
    coords1, _ := sliceAtoi(strings.Split(parts[0], ","))
    coords2, _ := sliceAtoi(strings.Split(parts[2], ","))

    return coord {
        x1: coords1[0],
        y1: coords1[1],
        x2: coords2[0],
        y2: coords2[1],
    }
}

func sliceAtoi(sa []string) ([]int, error) {
    si := make([]int, 0, len(sa))
    for _, a := range sa {
        i, err := strconv.Atoi(a)
        if err != nil {
            return si, err
        }
        si = append(si, i)
    }
    return si, nil
}

func createHVLine(c1, c2 int)(l []int) {
    if c1 > c2 {
        for ; c2 <= c1; c2++ {
            l = append(l, c2)
        }
    } else {
        for ; c1 <= c2; c1++ {
            l = append(l, c1)
        }
    }
    return
}

func createDiagLine (c coord) [][2]int {
    var points [][2]int
    grad := (c.y2 - c.y1) / (c.x2 - c.x1)
    i := c.y1 - grad * c.x1

    for x := c.x1; x <= c.x2; x++ {
        y := grad * x + i
        p := [2]int{x, y}
        points = append(points, p)
    }
    return points
}

