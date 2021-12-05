package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {

    data := readInput("./input.txt")
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

func readInput(filename string)  []string {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    lines := make([]string, 0)

    for sc.Scan() {
        lines = append(lines, sc.Text())
    }

    if err := sc.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}

func parseLines(s string) (string, int) {
    parts := strings.Fields(s)

    dist, err := strconv.Atoi(parts[1])
    if err != nil {
        log.Fatal(err)
    }
    return parts[0], dist
}
