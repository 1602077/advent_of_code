package main

import (
    "fmt"
    "strconv"
    "strings"
    "io/ioutil"
    "log"
)

func main() {
    part1()
}


func part1() int{
    data := readInputAsString("./input.txt")
    sects := strings.Split(data, "\n\n")

    numberDraw := parseNumberDraw(sects[0])
    bingoBoards := parseAllBingoBoards(sects[1:])

    fmt.Printf("Numbers called: %v.\n\n", numberDraw)
    fmt.Printf("Bingo Boards 1 & 2: \n %v \n\n %v\n\n", bingoBoards[0], bingoBoards[1])

    for _, num := range numberDraw {
        for _, b := range bingoBoards {
            if b.callNum(num) {
                fmt.Printf("Final Number: %v.\n", num)
                fmt.Printf("Final Board: %v.\n", b)
                fmt.Printf("Score of winning board: %v.\n", b.finalScore(num))
                return b.finalScore(num)
            }
        }
    }
    return -1
}


 /*
 ##################################################
 PARSE INPUT INTO NUMBERS DRAWN & BINGO BOARDS
 ##################################################
 */

func readInputAsString(filename string) string {
    content, err := ioutil.ReadFile("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    return string(content)
}


// Get called numbers from first line of `input.txt`
func parseNumberDraw(line string) []int {
    calledNumStr := strings.Split(line, ",")
    var numberDraw []int
    for _, value := range calledNumStr {
        num, _ := strconv.ParseInt(value, 10, 16)
        numberDraw = append(numberDraw, int(num))
    }
    return numberDraw
}

// extract a single bingo board from an input string
func parseBingoBoard(input string) *bingoBoard {
    b := bingoBoard{
        grid: make([]int, 25),
        called: make([]bool, 25),
    }
    for i, field := range strings.Fields(input) {
        num, _ := strconv.Atoi(field)
        b.grid[i] = int(num)
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

 /*
 ##################################################
 BINGO BOARD MANIPULATIONS
 ##################################################
 */
type bingoBoard struct {
    grid []int
    called []bool
}

 func (b *bingoBoard) lineFinished(line int) bool {
     return b.called[5*line] == true &&
        b.called[5*line + 1] == true &&
        b.called[5*line + 2] == true &&
        b.called[5*line + 3] == true &&
        b.called[5*line + 4] == true
 }

 func (b *bingoBoard) colFinished(col int) bool {
     return b.called[col] == true &&
        b.called[col + 5] == true &&
        b.called[col + 10] == true &&
        b.called[col + 15] == true &&
        b.called[col + 20] == true
 }

 func findIndex(el int, data []int) int {
     for k, v := range data {
         if el == v {
             return k
         }
     }
     return -1
}

func (b *bingoBoard) callNum(num int) bool {
    idx := findIndex(num, b.grid)
    if idx == -1 {
        // b.called[idx] = false
        return false
    }

    b.called[idx] = true

    return b.lineFinished(idx/5) || b.colFinished(idx%5)
 }

func (b *bingoBoard) finalScore(num int) int {
    sum := 0
    for i, val := range b.called {
        if val == false {
            sum += b.grid[i]
        }
    }
    return sum * num
}
