package main

import (
	"fmt"
	tempconv "gopl/ch2/2.1"
	"os"
	"strconv"
)

func CToF(t float64) {
	c := tempconv.Celsius(t)
	fmt.Printf("%.2f°C => %.2f°F\n", c, tempconv.CToF(c))
}

func FToC(t float64) {
	f := tempconv.Fahrenheit(t)
	fmt.Printf("%.2f°F => %.2f°C\n", f, tempconv.FToC(f))
}

func MToF(t float64) {
	fmt.Printf("%.2fm => %.2fft\n", t, t*3.28084)
}

func FToM(t float64) {
	fmt.Printf("%.2fft => %.2fm\n", t, t/3.28084)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			continue
		}

		CToF(t)
		FToC(t)
		MToF(t)
		FToM(t)

	}
}
