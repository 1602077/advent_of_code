package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)
func main() {
    part1("./input.txt")
}

func part1(filename string) int {
    playerPos := readInput(filename)

    var roll int = 0
    var scores = [2]int{0, 0}
    turn := 0
    numRolls := 0

    for scores[0] < 1000 && scores[1] < 1000 {
        for r := 0; r < 3; r++  {
            roll = roll % 100 + 1
            playerPos[turn] += roll
            numRolls++
        }
        playerPos[turn] = (playerPos[turn] - 1) % 10 + 1
        scores[turn] += playerPos[turn]
        turn = (turn + 1) % 2
    }

    minScore := 1000000
    for _, s := range scores {
        if  minScore > s {
            minScore = s
        }
    }
    fmt.Println("PART 1 Losing score * num rolls: %v.\n", minScore*numRolls)
    return minScore * numRolls
}

func readInput(filename string) [2]int {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }

    var pos [2]int
    line_str := strings.Split(string(data), "\n")
    for i, line := range line_str[:len(line_str)-1] {
        var tmp string
        fmt.Sscanf(line, "Player %v starting position: %d", &tmp, &pos[i])
    }
    return pos
}

