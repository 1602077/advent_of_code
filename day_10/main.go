package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main() {
    // part1("./input_test.txt")
    part1("./input.txt")
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
    // fmt.Printf("Input data:\n%v\n", input)

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
    fmt.Printf("PART 1 Total syntax error score: %v.\n", p)
    return p
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
