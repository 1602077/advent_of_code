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
    part2("./input.txt")
}

func part1(filename string) int {
     cP := ReadInput(filename)
    sort.Ints(cP)
    median := cP[len(cP)/2]
    // fmt.Printf("PART 1: Median distance for crabs to travel: %v.\n", median)

    fuel := 0
    for _, v := range cP {
        dist := v - median
        if dist < 0 {
            dist = -dist
        }
        fuel += dist
    }
    fmt.Printf("PART 1: Total Fuel required for all crabs to align: %v.\n", fuel)
    return fuel
}

func part2(filename string) int {
    cP := ReadInput(filename)

    targetPos := 0
    prevFuel := 0

    for {
        fuel := calcFuelP2(cP, targetPos)
        if fuel > prevFuel && prevFuel != 0 {
            break
        }
        targetPos += 1
        prevFuel = fuel
    }
    fmt.Printf("PART 2: Total fuel required for all crabs to align: %v.\n", prevFuel)
    return prevFuel
}

func calcFuelP2(pos []int, targetPos int) int {
    fuel := 0
    for _, p := range pos {
        f := 0
        dist := abs(targetPos - p)
        for d :=  1; d <= dist; d++ {
            f += d
        }
        fuel += f
    }
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

func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}
