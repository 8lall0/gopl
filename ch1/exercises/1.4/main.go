package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			err = f.Close()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v", err)
				return
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s: ", line)
			for _, filename := range filenames[line] {
				fmt.Printf("%s ", filename)
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileMap map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		if checkFileName(f.Name(), fileMap[input.Text()]) == true {
			fileMap[input.Text()] = append(fileMap[input.Text()], f.Name())
		}
	}
}

func checkFileName(filename string, fileList []string) bool {
	for _, name := range fileList {
		if filename == name {
			return false
		}
	}
	return true
}
