package main

import (
	"fmt"
	"gopl/ch2/exercises/2.1/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s = %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
		fmt.Printf("%s = %s = %s\n", c, tempconv.CToF(c), tempconv.CToK(c))
		fmt.Printf("%s = %s = %s\n", k, tempconv.KToC(k), tempconv.KToF(k))
	}
}
