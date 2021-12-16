package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "math"
    "strings"
)

func main() {
    part1("./input.txt")
    part2("./input.txt")
}

func part1(filename string) int {
    pT, mappings := readInput(filename)
    ptPairs := getPairs(pT)

    poly := string(ptPairs[0][0])
    final_poly := ""

    const NSTEPS = 10

    for k := 1; k <= NSTEPS; k++ {
        for _, p := range ptPairs {
            idx := findIndex(p, mappings.pair)
            poly += mappings.insert[idx]
        }
        // fmt.Printf("Step No. %v: Length of poly chain: %v.\n", k, len(poly))
        ptPairs = getPairs(poly)
        final_poly = poly
        poly = string(ptPairs[0][0])
    }

    type charCount struct {
        char []string
        count []int
    }

    var cc charCount

    for _, v := range final_poly {
        idx := findIndex(string(v), cc.char)
        if idx == -1 {
            cc.char = append(cc.char, string(v))
            cc.count = append(cc.count, 1)
        } else {
            cc.count[idx] += 1
        }
    }
    max, min := minMaxSlice(cc.count)
    fmt.Printf("PART 1 Max - min char count: %v.\n", max - min)
    return max - min

}

func part2(filename string) {
    pT, mappings := readInput(filename)
    allPairs := make(map[string]int)
    for i := 0; i < len(pT)-1; i++ {
        allPairs[string(pT[i])+string(pT[i+1])]++
    }

    const NSTEPS = 40

    for j := 0; j < NSTEPS; j++ {
        updatePairs := make(map[string]int)
        for k, v := range allPairs {
            idx := findIndex(k, mappings.pair)
            updatePairs[string(k[0]) + mappings.match[idx]] += v
            updatePairs[mappings.match[idx]+string(k[1])] += v
        }
        allPairs = updatePairs
    }

    counts := map[string]int{}
    for k, v := range allPairs {
        counts[string(k[0])] += v
    }

    max, min := 0, math.MaxInt
    for _, v := range counts {
        if min > v {
            min = v
        }
        if  max < v {
            max = v
        }
    }
    fmt.Printf("PART 2 Max - min char count: %v.", max - min)
}

type Mappings struct {
    pair []string
    match []string
    insert []string
}

func readInput(filename string) (string, Mappings) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    dS := strings.Split(string(data), "\n\n")
    poly_temp, map_dict := dS[0], dS[1]

    var mappings Mappings
    mR := strings.Split(map_dict, "\n")

    for i := range mR[:len(mR)-1] {
        line := strings.Split(string(mR[i]), " -> ")
        insert_str :=  string(line[1][0]) + string(line[0][1])
        mappings.pair = append(mappings.pair, line[0])
        mappings.match = append(mappings.match, line[1])
        mappings.insert = append(mappings.insert, insert_str)
    }
    return poly_temp, mappings
}

func getPairs(input string) []string {
    var pairs []string
    for i := range input[:len(input)-1] {
        pairs = append(pairs, string(input[i])+string(input[i+1]))
    }
    return pairs
}

func findIndex(el string, slice []string) int {
    for k, v := range slice {
        if el == v {
            return k
        }
    }
    return -1
}

func minMaxSlice(array []int) (int, int) {
    var min int = array[0]
    var max int = array[0]
    for _, v := range array{
        if max < v {
            max = v
        }
        if min > v {
            min = v
        }
    }
    return max, min
}
