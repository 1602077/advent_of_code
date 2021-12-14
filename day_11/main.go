package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main() {
    part1("./input_test.txt")
}

func part1(filename string) int {
    config := readInput(filename)
    fmt.Println(config)

    var num_flashes int = 0
    for s := 1; s < 100; s++ {
        var f int
        config, f = simulateFlash(config)
        num_flashes += f
    }
    fmt.Printf("PART 1 Number of flashes after 100 steps: %v.\n", num_flashes)
    return num_flashes
}

func simulateFlash(config [][]int) ([][]int, int) {
    // TODO: INCOMPLETE LOGIC HERE
    var num_flashes int = 0
    for i := range config {
        for j := range config[0] {
            c := config[i][j]

            if c == 10 {
                num_flashes++
                // TODO: Finish here; add recursion function around flashing neighbours
            } else {
                c++
            }
        }
    }

    return config, num_flashes
}

func readInput(filename string) [][]int {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    inStr := strings.Split(string(data), "\n")
    m:= make([][]int, len(inStr)-1)
    for i, val := range inStr[:len(inStr)-1] {
        a := make([]int, len(inStr[0]))
        for j, num := range val {
            a[j] = int(num - '0')
        }
        m[i] = a
    }
    return m
}
