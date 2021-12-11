package main

import (
    "fmt"
    "io/ioutil"
    "log"
    // "reflect"
    "strings"
)

type sigPattern struct {
    signal []string
    output []string
}

func main() {
    // part1("./input_test.txt")
    part1("./input.txt")
    part2("./input.txt")
}

func part1(filename string) int {
    displaySignal := readInput(filename)
    /* 1: c and f; 4: b, c, d, f; 7: a, c, f; 8: a,b,c,d,e,f,g */
    count := 0
    for _, sig := range displaySignal {
        for _, dig := range sig.output {
            length := len(dig)
            if length == 2 || length == 3 || length == 4 || length == 7 {
                count++
            }
        }
    }
    fmt.Printf("PART 1: Number of times a 1, 4, 7 or 8 are displayed: %v.\n", count)
    return count
}

func part2(filename string) int {
    displaySignal := readInput(filename)
    total := 0
    for _, sig := range displaySignal {
        for _, dig := range sig.output {
            switch {
            case len(dig) == 2:
                total += 1
            case len(dig) == 3:
                total += 7
            case len(dig) == 4:
                total += 4
            case len(dig) == 7:
                total += 8
            default:
                // TODO: Logic for cracking complex mappings
            }
        }
    }
    fmt.Printf("PART 2: Total of all output values: %v.\n", total)
    return total
}

func readInput(filename string) []sigPattern {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    var sigPattern []sigPattern
    lines := strings.Split(string(data), "\n")

    for _, line := range lines[:len(lines)-1] {
        dl := parseDisplayLine(line)
        sigPattern = append(sigPattern, dl)
    }
    return sigPattern
}

func parseDisplayLine(line string) sigPattern{
    line_str := strings.Split(line, " | ")
    var lineSignal sigPattern
    lineSignal.signal = strings.Split(strings.TrimSpace(line_str[0]), " ")
    lineSignal.output = strings.Split(strings.TrimSpace(line_str[1]), " ")
    return lineSignal
}
