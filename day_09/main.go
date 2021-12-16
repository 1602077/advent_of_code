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
    part2("./input.txt")
    // 1134
}

type grid struct {
    heightMap [][]int
    lpValue []int
    lpPos [][]int
}

type point struct{
    x, y int
}

var (
    dirs = []point{
        {x: -1, y: 0},
        {x: 0, y: -1},
        {x: 1, y: 0},
        {x: 0, y: 1},
    }
)

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

func part2(filename string) int {
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
    basinSizes := make([]int, 0)
    for _, p := range lT.lpPos {
        var pp point
        pp.x, pp.y  = p[0], p[1]
        points := calculateBasinSize(pp, lT.heightMap)
        basinSizes = append(basinSizes, len(points))
    }
    sort.Ints(basinSizes)
    l := len(basinSizes)
    prod := basinSizes[l-1]*basinSizes[l-2]*basinSizes[l-3]
    fmt.Printf("Part 2 Product of top three basin sizes: %v.\n", prod)
    return prod
}

func neighbors(p point, grid [][]int) []point {
    var result []point
    for _, d := range dirs {
        np := point{x: p.x + d.x, y: p.y + d.y}
        if np.x >= 0 && np.x < len(grid[p.y]) && np.y >= 0 && np.y < len(grid) {
            if grid[np.y][np.x] > grid[p.y][p.x] && grid[np.y][np.x] != 9 {
                result = append(result, np)
            }
        }
    }
    return result
}

func calculateBasinSize(p point, grid [][]int) map[point]struct{} {
    start := p
    seen := map[point]struct{}{
        start: {},
    }
    queue := []point{start}
    var current point
    for len(queue) > 0 {
        current, queue = queue[0], queue[1:]
        for _, next := range neighbors(current, grid) {
            if _, ok := seen[next]; !ok {
                queue = append(queue, next)
                seen[next] = struct{}{}
            }
        }
    }
    return seen
}

func inSlice(sub []int, slice [][]int) bool {
    for _, s := range slice {
        if len(s) != len(sub) {
            return false
        }
        for i, v := range s {
            if v != sub[i] {
                return false
            }
        }
    }
    return true
}

func isLowPoint(matrix [][]int, i, j int) bool {
    c := 0
    el := matrix[i][j]
    if i-1 < 0 || matrix[i-1][j] > el {
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
