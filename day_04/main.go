package main

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/1602077/advent_of_code/utils"
)

func main() {
    splitInputs("./input.txt")

}

type BingBoard struct {
    grid [][]int
    called [][]bool
}

func splitInputs(filename string){
    data := utils.ReadInput(filename)
    numBoards := (len(data) - 1) / 6

    fmt.Printf("Number of lines in dataset: %v.\n", len(data))
    fmt.Printf("Number of boards in dataset: %v.\n", numBoards)

    numberDraw := parseNumberDraw(data[0])

    bingoGrids := []BingBoard{}

}

func parseNumberDraw(line string) []int16 {
    calledNumStr := strings.Split(line, ",")
    var numberDraw []int16
    for _, value := range calledNumStr {
        num, _ := strconv.ParseInt(value, 10, 16)
        numberDraw = append(numberDraw, int16(num))
    }
    fmt.Printf("Numbers called: %v.\n", numberDraw)
    return numberDraw
}
