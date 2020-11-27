package main

import (
	"fmt"
	"strings"
)

func removeDup(s []string) []string {
	w := 0
	for _, c := range s {
		if s[w] == c {
			continue
		}
		w++
		s[w] = c
	}

	return s[:w+1]
}

func main() {
	s := strings.Split("aaaantonio", "")
	s = removeDup(s)
	fmt.Println(s)
}
