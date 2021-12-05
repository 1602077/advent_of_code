package utils

import (
    "log"
    "os"
    "bufio"
)

func readInput(filename string)  []string {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    lines := make([]string, 0)

    for sc.Scan() {
        lines = append(lines, sc.Text())
    }

    if err := sc.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}
