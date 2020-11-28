package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

func main() {
	freq, _ := tagCount(os.Stdin)
	for k, v := range freq {
		fmt.Printf("[%s]: %d\n", k, v)
	}
}

func tagCount(r io.Reader) (map[string]int, error) {
	tags := make(map[string]int)
	z := html.NewTokenizer(r)

	var err error

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken {
			name, _ := z.TagName()
			tags[string(name)]++
		}
	}

	if err != io.EOF {
		return tags, err
	}

	return tags, nil
}
