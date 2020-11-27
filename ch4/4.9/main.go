package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	fIn, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Charcount: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(fIn)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	fmt.Printf("Key           Value\n")
	for k, v := range counts {
		fmt.Printf("%10s\t%d\n", k, v)
	}

}
