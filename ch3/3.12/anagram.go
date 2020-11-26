package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isAnagram(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}

	_s1 := []rune(s1)
	_s2 := []rune(s2)
	sort.Sort(sortRunes(_s1))
	sort.Sort(sortRunes(_s2))

	return string(_s1) == string(_s2)
}

func main() {
	fmt.Println(isAnagram("coloc", "coasl"))
	fmt.Println(isAnagram("coloc", "coloc"))
	fmt.Println(isAnagram("coloc", "cocol"))
}
