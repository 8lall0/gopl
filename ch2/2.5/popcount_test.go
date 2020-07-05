package popcount

import "testing"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>0*8)] +
		pc[byte(x>>1*8)] +
		pc[byte(x>>2*8)] +
		pc[byte(x>>3*8)] +
		pc[byte(x>>4*8)] +
		pc[byte(x>>5*8)] +
		pc[byte(x>>6*8)] +
		pc[byte(x>>7*8)])
}

func PopCountLoop(x uint64) int {
	j := 0
	for i := 0; i < 8; i++ {
		j += int(pc[byte(x>>i*8)])
	}

	return j
}

func PopCountShift(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCountClearRightmost(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkPopCount(t *testing.B) {
	bench(t, PopCount)
}

func BenchmarkPopCountLoop(t *testing.B) {
	bench(t, PopCountLoop)
}

func BenchmarkPopCountShift(t *testing.B) {
	bench(t, PopCountShift)
}

func BenchmarkClearRightmost(t *testing.B) {
	bench(t, PopCountClearRightmost)
}
