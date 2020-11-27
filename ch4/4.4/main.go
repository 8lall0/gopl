package main

import "fmt"

func rotate(s []int) {
	tmp := s[0]
	copy(s, s[1:])
	s[len(s)-1] = tmp
}

func main() {
	var a = []int{0, 1, 2, 3, 4}
	rotate(a)
	fmt.Println(a)
}
