package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	Control = iota
	Letter
	Mark
	Number
	Punct
	Space
	Symbol
	Graphic
	Print
	UTFCatCount
)

func main() {
	counts := make(map[rune]int)
	var utfLen [utf8.UTFMax + 1]int
	var utfcat [UTFCatCount]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsPrint(r) {
			utfcat[Print]++
		}

		if unicode.IsGraphic(r) {
			utfcat[Graphic]++
		}

		switch {
		case unicode.IsControl(r):
			utfcat[Control]++
		case unicode.IsLetter(r):
			utfcat[Letter]++
		case unicode.IsMark(r):
			utfcat[Mark]++
		case unicode.IsNumber(r):
			utfcat[Number]++
		case unicode.IsPunct(r):
			utfcat[Punct]++
		case unicode.IsSymbol(r):
			utfcat[Symbol]++
		case unicode.IsSpace(r):
			utfcat[Space]++
		}

		counts[r]++
		utfLen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("rune\tcount\n")
	for i, n := range utfLen {
		if i > 0 {
			fmt.Printf("%q\t%d\n", i, n)
		}
	}

	fmt.Print("\nCategory   Count\n")
	fmt.Printf("%-7.7s: %4d\n", "Print", utfcat[Print])
	fmt.Printf("%-7.7s: %4d\n", "Graphic", utfcat[Graphic])
	fmt.Printf("%-7.7s: %4d\n", "Control", utfcat[Control])
	fmt.Printf("%-7.7s: %4d\n", "Letter", utfcat[Letter])
	fmt.Printf("%-7.7s: %4d\n", "Mark", utfcat[Mark])
	fmt.Printf("%-7.7s: %4d\n", "Number", utfcat[Number])
	fmt.Printf("%-7.7s: %4d\n", "Punct", utfcat[Punct])
	fmt.Printf("%-7.7s: %4d\n", "Space", utfcat[Space])
	fmt.Printf("%-7.7s: %4d\n", "Symbol", utfcat[Symbol])

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters.\n", invalid)
	}
}
