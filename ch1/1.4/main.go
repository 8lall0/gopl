package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, foundIn)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, foundIn)
			_ = f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, foundIn[line])
		}
	}
}

func isFileIn(name string, array []string) bool {
	for _, s := range array {
		if name == s {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !isFileIn(line, foundIn[line]) {
			foundIn[line] = append(foundIn[line], f.Name())
		}
	}
}
