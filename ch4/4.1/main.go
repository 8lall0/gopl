package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Printf("Diff: %d", compareSha256([]byte("a"), []byte("b")))
}

func compareSha256(a, b []byte) int {
	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)

	return diff(shaA[:], shaB[:])
}

func popCount(b byte) int {
	count := 0

	for ; b != 0; count++ {
		b &= b - 1
	}

	return count
}

func diff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i])
		}
	}

	return count
}
