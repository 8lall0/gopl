package testing

import (
	"strings"
	"testing"
)

var args = []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

func echoSum() {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

func echoJoin() {
	strings.Join(args, " ")
}

func BenchmarkEchoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoSum()
	}
}

func BenchmarkEchoJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoJoin()
	}
}
