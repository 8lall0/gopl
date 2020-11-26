package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(comma(s))
	}
}

func comma(s string) string {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return ""
	}
	var buf bytes.Buffer

	mantissaStart := 0
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteString(s[0:1])
		mantissaStart = 1
	}

	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}

	mantissa := s[mantissaStart:mantissaEnd]

	start := len(mantissa) % 3
	if start == 0 {
		start = 3
	}
	buf.WriteString(mantissa[0:start])
	for i := start; i < len(mantissa); i += 3 {
		buf.WriteString(",")
		buf.WriteString(mantissa[i : i+3])
	}

	buf.WriteString(s[mantissaEnd:])

	return buf.String()
}
