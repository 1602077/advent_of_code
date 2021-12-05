package main

import (
    "fmt"
    "log"
    "strconv"

    "github.com/1602077/advent_of_code/utils"
)

func main() {
    data := utils.ReadInput("./input.txt")

    solve1(data)
    solve2(data)
}

func solve1(data []string) {
    gamma, eps := "", ""
    for i := 0; i < len(data[0]); i++ {
        most, least := mostLeastCommonBit(data, i)
        gamma += string(most)
        eps += string(least)
    }
    fmt.Printf("Gamma: %v.\n", gamma)
    fmt.Printf("Epsilon: %v.\n", eps)
    fmt.Printf("Power Consumption: %v.\n", binStringToInt(gamma) * binStringToInt(eps))
}

func solve2(data []string) {
    oxy := filterBits(data, 0, true)
    scrub := filterBits(data, 0, false)
    fmt.Printf("Life Support Rating: %v.", binStringToInt(oxy) * binStringToInt(scrub))
}

func mostLeastCommonBit(lines []string, charPos int) (most uint8, least uint8) {
    count0, count1 := 0, 0
    for _, line := range lines {
        if line[charPos] == '0' {
            count0++
        } else {
            count1++
        }
    }

    if count0 > count1 {
        most, least = '0', '1'
    } else {
        most, least = '1', '0'
    }
    return most, least
}

func binStringToInt(s string)(int){
    parse, err := strconv.ParseInt(s, 2, 64)
    if err != nil {
        log.Fatal(err)
    }
    return int(parse)
}

func filterBits(lines []string, cn int, bitCriteria bool) string {
    if len(lines) == 1 {
        return lines[0]
    }
    most, least := mostLeastCommonBit(lines, cn)
    comparator := least
    if bitCriteria {
        comparator = most
    }
    filtered := make([]string, 0)
    for _, l := range lines {
        if l[cn] == comparator {
            filtered = append(filtered, l)
        }
    }
    return filterBits(filtered, cn+1, bitCriteria)

}
