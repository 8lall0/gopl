package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var width = flag.Int("w", 256, "Hash width (256, 384, 512")

func main() {
	flag.Parse()

	var function func(a []byte) []byte

	switch *width {
	case 256:
		function = func(a []byte) []byte {
			h := sha256.Sum256(a)
			return h[:]
		}
	case 384:
		function = func(a []byte) []byte {
			h := sha512.Sum384(a)
			return h[:]
		}
	case 512:
		function = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("Unexpected size")
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", function(b))

}
