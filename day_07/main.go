package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "sort"
    "strconv"
    "strings"
)
func main() {
    part1("./input.txt")
}

func part1(filename string) int {
     cP := ReadInput(filename)
    sort.Ints(cP)
    median := cP[len(cP)/2]
    fmt.Printf("Median distance for crabs to travel: %v.\n", median)

    fuel := 0
    for _, v := range cP {
        dist := v - median
        if dist < 0 {
            dist = -dist
        }
        fuel += dist
    }
    fmt.Printf("Total Fuel required for all crabs to align: %v.\n", fuel)
    return fuel
}

func ReadInput(filename string)  []int {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    var crabPos []int
    pos, err :=  sliceAtoi(strings.Split(strings.TrimSuffix(string(data), "\n"), ","))
    if err != nil {
        log.Fatal(err)
    }

    for _, p := range pos[:len(pos)] {
        crabPos = append(crabPos, p)
    }
    return crabPos
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
