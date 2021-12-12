package main

import (
    // "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    // "strconv"
    // "reflect"
)

func main() {
    part1("./input.txt")
}

type lavaTube struct {
    heightMap [][]int
    lowPoints []int
}

func part1(filename string) int {
    m := readInput(filename)

    lT := lavaTube{heightMap: m, lowPoints: make([]int, 0)}

    nRows := len(lT.heightMap)
    nCols := len(lT.heightMap[0])

    r := 0
    for i := range lT.heightMap {
        for j := range lT.heightMap[0] {
            if isLowPoint(lT.heightMap, i, j) {
                lT.lowPoints = append(lT.lowPoints, lT.heightMap[i][j])
                r += 1 + int(lT.heightMap[i][j])
            }
        }
    }
    // Input file summary
    // fmt.Printf("Input heightmap: \n%v.\n", lT.heightMap)
    // fmt.Printf("Number of rows and columns: %v, %v.\n", nRows, nCols)
    // fmt.Printf("Lavatube low points: %v.\n", lT.lowPoints)

    fmt.Printf("PART 1 Total risk score: %v.\n", r)
    return r
}

func isLowPoint(matrix [][]int, i, j int) bool {
    c := 0
    el := matrix[i][j]
    if i-1 <0 || matrix[i-1][j] > el {
        c++
    }
    if i+1 >= len(matrix) || matrix[i+1][j] > el {
        c++
    }
    if j-1 < 0 || matrix[i][j-1] > el {
        c++
    }
    if j+1 >= len(matrix[0]) || matrix[i][j+1] > el {
        c++
    }
    return c == 4
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
