package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma(s))
	}

}

func comma(s string) string {
	if _, err := strconv.Atoi(s); err != nil {
		return ""
	}

	var buf bytes.Buffer
	start := len(s) % 3
	if start == 0 {
		start = 3
	}
	buf.WriteString(s[0:start])

	for i := start; i < len(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
