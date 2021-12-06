package main

import (
    "fmt"
    "strconv"
    "strings"
    "io/ioutil"
    "log"

    // "github.com/1602077/advent_of_code/utils"
)

func main() {
    splitInputs("./input.txt")

}

type bingoBoard struct {
    grid []byte
    called [][]bool
}

func readInputAsString(filename string) string {
    content, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}

func splitInputs(filename string){
    data := readInputAsString(filename)
    sects := strings.Split(data, "\n\n")

    numberDraw := parseNumberDraw(sects[0])
    bingoBoards := parseAllBingoBoards(sects[1:])

    fmt.Printf("Numbers called: %v.\n\n", numberDraw)
    fmt.Printf("Bingo Boards 1 & 2: \n %v \n\n %v", bingoBoards[0], bingoBoards[1])
}

// Get called numbers from first line of `input.txt`
func parseNumberDraw(line string) []int16 {
    calledNumStr := strings.Split(line, ",")
    var numberDraw []int16
    for _, value := range calledNumStr {
        num, _ := strconv.ParseInt(value, 10, 16)
        numberDraw = append(numberDraw, int16(num))
    }
    return numberDraw
}

// extract a single bingo board from an input string
func parseBingoBoard(input string) *bingoBoard {
    b := bingoBoard{
        grid: make([]byte, 25),
    }
    for i, field := range strings.Fields(input) {
        num, _ := strconv.Atoi(field)
        b.grid[i] = byte(num)
    }
    return &b
}

// Parse over all bingo boards in the dataset and return as a slice
func parseAllBingoBoards(input []string) []bingoBoard {
    var bingoBoards []bingoBoard
    for _, part := range input {
        bingoBoards = append(bingoBoards, *parseBingoBoard(part))
    }
    return bingoBoards
}
