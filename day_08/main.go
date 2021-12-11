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
}

func part1(filename string) int {
    displaySignal := readInput(filename)

    // count how many times 1, 4, 7, 8 appear in .output
    count := 0
    /*
    1: c and f on
    4: b, c, d, f
    7: a, c, f
    8: a,b,c,d,e,f,g
    */
    for _, sig := range displaySignal {
        fmt.Println(sig.output)
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
