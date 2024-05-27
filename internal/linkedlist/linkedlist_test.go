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

func Example_pSingleSearch() {
	ll := NewPSingleLinkedList[int]()
	for i := 10; i < 20; i++ {
		ll.Append(i)
	}
	n := ll.Search(0)
	fmt.Printf("Node ID: %v Value: %v\n", n.ID, n.Value)
	n = ll.Search(3)
	fmt.Printf("Node ID: %v Value: %v\n", n.ID, n.Value)
	n = ll.Search(10)
	if n == nil {
		fmt.Println("Out of range")
	}

	// Output:
}

func BenchmarkPSingleLLSearch(b *testing.B) {

	benchcases := []struct {
		input uint
	}{
		{input: 0},
		{input: 6},
		{input: 15},
		{input: 30},
		{input: 40},
		{input: 49},
		{input: 50},
	}

	for _, bench := range benchcases {
		b.Run(fmt.Sprintf("input_%d", bench.input), func(b *testing.B) {
			ll := NewPSingleLinkedList[int]()
			for i := 10; i < 50; i++ {
				ll.Append(i)
			}
			for count := 0; count < b.N; count++ {
				ll.Search(bench.input)
			}
		})
	}
}
