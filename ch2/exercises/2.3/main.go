package main

import (
	"fmt"
	"gopl/ch2/exercises/2.3/popcount"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
	parsed := uint64(n)
	fmt.Printf("PopcountSum: %d\n", popcount.PopcountSum(parsed))
	fmt.Printf("PopcountLoop: %d\n", popcount.PopcountLoop(parsed))
}
