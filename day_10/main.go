package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "sort"
    "strings"
)

func main() {
    part1("./input.txt")
    part2("./input.txt")
}

type Stack []byte

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

func (s *Stack) Push(str byte) {
    *s = append(*s, str)
}

func (s *Stack) Pop() (byte, bool) {
    if s.IsEmpty() {
        return ' ', false
    } else {
        idx := len(*s) - 1
        el := (*s)[idx]
        *s = (*s)[:idx]
        return el, true
    }
}

func part1(filename string) int {
    input := readInput(filename)

    points := map[byte]int{
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137,
    }

    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
        '>': '<',
    }
    p := 0
    for _, line := range input {
        if ok, l := isIllegal(line, pairs); ok {
            p += points[l]
        }
    }
    fmt.Printf("PART 1 Total syntax error score: %v\n", p)
    return p
}

func part2(filename string) int {
    input := readInput(filename)

    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
        '>': '<',
    }

    var incomplete_lines []string

    for _, line := range input {
        corrupt, _ :=  isIllegal(line, pairs)
        if corrupt == false {
            incomplete_lines = append(incomplete_lines, line)
        }
    }

    cScore := map[byte]int{
        '(': 1,
        '[': 2,
        '{': 3,
        '<': 4,
    }

    scores := make([]int, len(incomplete_lines))
    for i, str := range incomplete_lines {
        r := incomplete(str)
        p := 0
        for !r.IsEmpty() {
            v, _ := r.Pop()
            p = 5*p + cScore[v]
        }
        scores[i] = p
    }

    sort.Ints(scores)
    fmt.Printf("PART 2 Total score to complete all lines: %v.\n", scores[len(scores)/2])
    return scores[len(scores)/2]
}

func incomplete(line string) Stack {
    var stack Stack
    for i := range line {
        if isOpen(line[i]) {
            stack.Push(line[i])
        } else {
            stack.Pop()
        }
    }
    return stack
}

func isOpen(str byte) bool {
    return str == '(' || str == '[' || str == '{' || str == '<'
}

func isIllegal(line string, pairs map[byte]byte) (bool, byte) {
    var stack Stack
    for i := range line {
        if isOpen(line[i]) {
            stack.Push(line[i])
        } else {
            top, _ := stack.Pop()
            if pairs[line[i]] != top {
                return true, line[i]
            }
        }
    }
    return false, ' '
}

func readInput(filename string) []string {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    return strings.Split(string(data), "\n")
}
