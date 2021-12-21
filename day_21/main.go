package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)
func main() {
    part1("./input.txt")
    part2("./input.txt")
}

func part1(filename string) int {
    playerPos := readInput(filename)

    var roll int = 0
    var scores = []int{0, 0}
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
    fmt.Printf("PART 1 Losing score * num rolls: %v.\n", minOfSlice(scores)*numRolls)
    return minOfSlice(scores) * numRolls
}

type player struct {
    pos int
    score int
}

type game struct {
    p1, p2 player
}

func part2(filename string) {
    p := readInput(filename)
    p1 := player{pos: p[0]}
    p2 := player{pos: p[1]}

    win := playGame(game{p1: p1, p2: p2})
    fmt.Printf("PART 2 Player 1 win count: %v.\n", win[0])
    fmt.Printf("PART 2 Player 2 win count: %v.\n", win[1])
}

var dp = make(map[game][]int)

func playGame(g game) []int {
    if g.p1.score >= 21 {
        return []int{1, 0}
    }
    if g.p2.score >= 21 {
        return []int{0, 1}
    }
    if v, ok := dp[g]; ok {
        return v
    }
    win := []int{0, 0}
    for d1 := 1; d1 <= 3 ; d1++ {
        for d2 := 1 ; d2 <= 3; d2++ {
            for d3 := 1; d3 <= 3; d3++ {
                p1 := (g.p1.pos + d1 + d2 + d3 - 1) % 10 +1
                s1 := g.p1.score + p1

                w := playGame(game{
                    p1: player{pos: g.p2.pos, score: g.p2.score},
                    p2: player{pos: p1, score: s1},
                })

                win[0] += w[1]
                win[1] += w[0]
            }
        }
    }
    dp[g] = win
    return win
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

func minOfSlice(slice []int) int {
    minScore := 100000000000000000
    for _, s := range slice {
        if minScore > s {
            minScore = s
        }
    }
    return minScore
}

