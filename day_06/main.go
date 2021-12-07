package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
)

func main() {
    solve(80)
}

func solve(numDays int) {
    const (
        internalTimer = 8
        minTimer = 0
    )

    lfs := parseInput()

    newLfsC := 0
    for k:=0; k < numDays; k++ {
        for i := range lfs {
            lfs[i]--
            if lfs[i] == -1 {
                lfs[i] = 6
                newLfsC++
            }
        }
        for j := 1; j <=  newLfsC; j++ {
            if newLfsC == 0 {
                break
            }
            lfs = append(lfs, internalTimer)
        }
        newLfsC = 0
    }
    fmt.Printf("Number of Lantern fish after %v days: %v.\n", numDays, len(lfs))
}

func parseInput() []int {
    data, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        log.Fatal(err)

    }
    lfs_string := strings.Split(strings.TrimSuffix(string(data), "\n"), ",")
    lfs, err := sliceAtoi(lfs_string)
    if err != nil {
        log.Fatal(err)
    }
    return lfs
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
