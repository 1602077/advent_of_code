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
}

func part1() {
    parseInputFile()
}


type coord struct {
    x1 int
    y1 int
    x2 int
    y2 int
}

func parseInputFile() {
    data, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        log.Fatal(err)
    }

    lineSegs := []coord

    // coords := parseCoord("0,9 -> 5,9")
    // fmt.Println(coords)

}

func parseCoord(input string) coord {
    parts := strings.Fields(input)
    // Convert []str -> []int.
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

