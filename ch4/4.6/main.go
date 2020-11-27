package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpace(s []byte) []byte {
	out := s[:0]
	var last rune

	for i := 0; i < len(s); {
		r, space := utf8.DecodeRune(s[i:])

		if !unicode.IsSpace(r) {
			out = append(out, s[i:i+space]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			out = append(out, ' ')
		}

		last = r
		i += space
	}

	return out
}

func main() {
	s := []byte("Try    try try    try")
	fmt.Println(string(squashSpace(s)))
}
