package linkedlist

import (
	"fmt"
	"testing"
)

func Example_pSingleLL() {

	ll := NewPSingleLinkedList[int]()
	ll.Append(1)
	ll.Append(2)

	v := ll.Values()
	fmt.Println(v)

	// Output:
	// [1 2]
}

func BenchmarkPSingleLLAppend(b *testing.B) {

	benchcases := []struct {
		input int
	}{
		{input: 1},
		{input: 10},
		{input: 20},
		{input: 30},
	}

	for _, bench := range benchcases {
		b.Run(fmt.Sprintf("input_%d", bench.input), func(b *testing.B) {
			ll := NewPSingleLinkedList[int]()
			for count := 0; count < b.N; count++ {
				ll.Append(bench.input)
			}
		})
	}
}
