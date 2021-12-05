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
}

func solve1(data []string) {
    gamma, eps := "", ""
    for i := 0; i < len(data[0]); i++ {
        most, least := mostLeastCommonBit(data, i)
        gamma += string(most)
        eps += string(least)
    }
    fmt.Printf("Gamma: %v.\n", gamma)
    fmt.Printf("Epsilon: %v.\n\n", eps)

    fmt.Printf("Power Consumption: %v.\n", binStringToInt(gamma) * binStringToInt(eps))
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
