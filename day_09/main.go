package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "sort"
    // "strconv"
)

func main() {
    part1("./input.txt")
    part2("./input_test.txt")
}

type grid struct {
    heightMap [][]int
    lpValue []int
    lpPos [][]int
}

func part1(filename string) int {
    m := readInput(filename)
    lT := grid{heightMap: m, lpValue: make([]int, 0), lpPos: make([][]int, 0)}

    r := 0
    for i := range lT.heightMap {
        for j := range lT.heightMap[0] {
            if isLowPoint(lT.heightMap, i, j) {
                lT.lpValue = append(lT.lpValue, lT.heightMap[i][j])
                lT.lpPos = append(lT.lpPos, []int{i,j})
                r += 1 + int(lT.heightMap[i][j])
            }
        }
    }
    fmt.Printf("PART 1 Total risk score: %v.\n", r)
    return r
}

func part2(filename string) {
    m := readInput(filename)
    lT := grid{heightMap: m, lpValue: make([]int, 0), lpPos: make([][]int, 0)}

    for i := range lT.heightMap {
        for j := range lT.heightMap[0] {
            if isLowPoint(lT.heightMap, i, j) {
                lT.lpValue = append(lT.lpValue, lT.heightMap[i][j])
                lT.lpPos = append(lT.lpPos, []int{i,j})
            }
        }
    }
    // TODO: finish basin size caclulation function
    getBasinSizes(lT)

}

func getBasinSizes(lt grid) []int {
    fmt.Println(lt.heightMap)
    fmt.Println(lt.lpValue)

    // nRows := len(lt.heightMap)
    // nCols := len(lt.heightMap[0])

    var basinSizes []int
    var inBasin [][]bool
    for k := range lt.lpValue {
        p := lt.lpPos[k]
        fmt.Println(inBasin[p[0]][p[1]])
        fmt.Print()
        break
    }

    /*
    for k, _ := range lt.lpValue {
        pos := lt.lpPos[k]
        fmt.Println(pos)
        var basinSize int = 1
        // iterate right
        for i := pos[0]+1; i < nRows; i++ {
            prev := lt.heightMap[i-1][pos[1]]
            curr := lt.heightMap[i][pos[1]]
            if curr > prev && curr != 9 {
                basinSize++
            }
        }
        // iterate left
        for i := pos[0]-1; i >= 0; i-- {
            prev := lt.heightMap[i+1][pos[1]]
            curr := lt.heightMap[i][pos[1]]
            if curr > prev && curr != 9 {
                basinSize++
            }
        }
        // iterate top
        for j := pos[1] + 1; j < nCols; j++ {
            prev := lt.heightMap[pos[0]][j-1]
            curr := lt.heightMap[pos[0]][j]
            if curr > prev && curr != 9 {
                basinSize++
            }
        }
        // iterate down
        for j := pos[1] - 1; j >= 0; j-- {
            prev := lt.heightMap[pos[0]][j+1]
            curr := lt.heightMap[pos[0]][j]
            if curr > prev && curr != 9 {
                basinSize++
            }
        }
        basinSizes = append(basinSizes, basinSize)
    }
    */
    fmt.Printf("PART 2 Basin Sizes: %v.\n", basinSizes)

    sort.Ints(basinSizes)
    var prod int = 1
    for _, v := range basinSizes[len(basinSizes) - 3:] {
        prod *= v
    }
    fmt.Printf("Part 2 Product of basin sizes: %v.\n", prod)
    return basinSizes
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
