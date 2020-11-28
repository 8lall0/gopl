package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

func main() {
	printTxt(os.Stdin, os.Stdout)
}

func printTxt(r io.Reader, w io.Writer) {
	z := html.NewTokenizer(r)
	avoid := false

	for {
		switch z.Next() {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			b, _ := z.TagName()
			if string(b) == "script" || string(b) == "style" {
				avoid = true
			}
		case html.TextToken:
			if avoid == true {
				continue
			}
			txt := z.Text()
			if len(strings.TrimSpace(string(txt))) == 0 {
				continue
			}
			_, _ = fmt.Fprintf(w, "%s", txt)
		case html.EndTagToken:
			b, _ := z.TagName()
			if string(b) == "script" || string(b) == "style" {
				avoid = false
			}
		}
	}
}
