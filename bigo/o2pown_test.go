package bigo

import (
	"fmt"
	"testing"
)

// operation to count to zero
func countToZero(number int) int {
	if number < 1 {
		return number
	}
	number = number - 1
	return countToZero(number)
}

// O(2^N)
var counttests = []struct {
	input int
}{
	{
		input: 1,
	},
	{
		input: 50,
	},
	{
		input: 100,
	},
	{
		input: 1_000,
	},
}

func Benchmark2PowN(b *testing.B) {
	for _, test := range counttests {
		b.Run(fmt.Sprintf("input_%d", test.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				countToZero(test.input)
			}
		})
	}
}
