package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// part1(80)
	part2(80)
}

func part1(numDays int) {
	const (
		internalTimer = 8
		minTimer      = 0
	)

	lfs := parseInput()

	newLfsC := 0
	for k := 0; k < numDays; k++ {
		for i := range lfs {
			lfs[i]--
			if lfs[i] == -1 {
				lfs[i] = 6
				newLfsC++
			}
		}
		for j := 1; j <= newLfsC; j++ {
			if newLfsC == 0 {
				break
			}
			lfs = append(lfs, internalTimer)
		}
		newLfsC = 0
	}
	fmt.Printf("\nNumber of Lantern fish after %v days: %v.\n", numDays, len(lfs))
}

func part2(numDays int) {
	const (
		internalTimer = 8
		minTimer      = 0
	)
	lfs := parseInput()

	// count the number of laternfish at each point in their internal replication cycle
	var lf_lifecycles [internalTimer + 1]int
	for _, lf := range lfs {
		lf_lifecycles[lf]++
	}
	fmt.Print(lf_lifecycles)

	var y, x int
	for k := 0; k < numDays; k++ {
		for i := internalTimer; i >= minTimer; i-- {
			if i == 0 {
				lf_lifecycles[internalTimer] += lf_lifecycles[i]
				lf_lifecycles[6] += lf_lifecycles[i]
			}
			x = lf_lifecycles[i]
			lf_lifecycles[i] = y
			y = x
		}
	}
	numLF := 0
	for i := internalTimer; i >= minTimer; i-- {
		numLF += lf_lifecycles[i]
	}
	fmt.Printf("\nNumber of Lantern fish after %v days: %v.\n", numDays, numLF)
}

func parseInput() []int {
	// data, err := ioutil.ReadFile("./input.txt")
	data, err := ioutil.ReadFile("./input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	lfs_string := strings.Split(strings.TrimSuffix(string(data), "\n"), ",")
	lfs, err := sliceAtoi(lfs_string)
	if err != nil {
		log.Fatal(err)
	}
	return lfs
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
