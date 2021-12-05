package main

import (
    "fmt"
    "log"
    "strconv"
    "strings"

    "github.com/1602077/advent_of_code/utils"
)

func main() {

    data := utils.ReadInput("./input.txt")
    fmt.Println("Part One Solution")
    solve1(data)
    fmt.Println("\nPart Two Solution")
    solve2(data)
}

func solve2(data []string) {
    x, y, aim := 0, 0, 0
    for _, line := range data {
        dir, dist := parseLines(line)

        switch dir {
        case "forward":
            x += dist
            y += aim * dist
        case "down":
            aim += dist
        case "up":
            aim -= dist
        }
    }

    fmt.Printf("x, y, aim: (%v, %v, %v).\nx*y: %v.\n", x, y, aim,  x*y)

}

func solve1(data []string) {
    x, y := 0, 0
    for _, line := range data {
        dir, dist := parseLines(line)

        switch dir {
        case "forward":
            x += dist
        case "down":
            y += dist
        case "up":
            y -= dist
        }
    }

    fmt.Printf("x, y: (%v, %v).\nx*y: %v.\n", x, y,  x*y)

}



func parseLines(s string) (string, int) {
    parts := strings.Fields(s)

    dist, err := strconv.Atoi(parts[1])
    if err != nil {
        log.Fatal(err)
    }
    return parts[0], dist
}
