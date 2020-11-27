package main

import "fmt"

func reverse(s *[5]int) {
	for i := 0; i < len(s)/2; i++ {
		end := len(s) - i - 1
		s[i], s[end] = s[end], s[i]
	}
}

func main() {
	var a = [5]int{0, 1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}
