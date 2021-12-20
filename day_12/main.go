package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

type cave struct {
    value string
    currPath []string
    twice bool
}

func main() {
    part1("./input.txt")
    part2("./input.txt")
}

func part1(filename string) int {
    caves := readInput(filename)

    count := 0
    start := cave{
        value: "start",
        currPath: []string{"start"},
    }
    queue := []cave{start}

    var curr cave

    for len(queue) > 0 {
        curr, queue = queue[0], queue[1:]
        if curr.value == "end" {
            count ++
            continue
        }
        for _, next := range caves[curr.value] {
            if !seenCave(next, curr.currPath) {
                path := make([]string, 0)
                path = append(path, curr.currPath...)
                if strings.ToLower(next) == next {
                    path = append(path, next)
                }
                queue = append(queue, cave{
                    value: next,
                    currPath: path,
                })
            }
        }
    }
    fmt.Printf("PART 1 Number of paths visiting just a small cave once: %v.\n", count)
    return count
}

func part2(filename string) int {
    caves := readInput(filename)

    count := 0
    start := cave{
        value: "start",
        currPath: []string{"start"},
    }
    queue := []cave{start}

    var curr cave

    for len(queue) > 0 {
        curr, queue = queue[0], queue[1:]
        if curr.value == "end" {
            count ++
            continue
        }
        for _, next := range caves[curr.value] {
            if !seenCave(next, curr.currPath) {
                path := make([]string, 0)
                path = append(path, curr.currPath...)
                if strings.ToLower(next) == next {
                    path = append(path, next)
                }
                queue = append(queue, cave{
                    value: next,
                    currPath: path,
                    twice: curr.twice,
                })
            } else if seenCave(next, curr.currPath) && !curr.twice && !seenCave(next, []string{"start", "end"}) {
                queue = append(queue, cave{
                    value: next,
                    currPath: curr.currPath,
                    twice: true,
                })
            }
        }
    }
    fmt.Printf("PART 2 Number of paths visiting just a single small cave twice: %v.\n", count)
    return count
}

func readInput(filename string) map[string][]string{
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    input := strings.Split(string(data), "\n")
    caves := make(map[string][]string)

    for _, line := range input[:len(input)-1] {
        split := strings.Split(line, "-")
        var a string = split[0]
        var b string = split[1]
        caves[a] = append(caves[a], b)
        caves[b] = append(caves[b], a)
    }
    return caves
}

func seenCave(el string, slice []string) bool {
    for _, v := range slice {
        if v == el {
            return true
        }
    }
    return false
}
