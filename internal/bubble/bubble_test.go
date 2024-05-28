package bubble

import (
	"fmt"
	"testing"
)

func Example_sortInt() {
	ds := Array[int]{-4, 1, 2, 5, 7, 3, -6, 10, 6}
	result := Sort(ds)
	fmt.Println(result)
	// Output:
	// [-6 -4 1 2 3 5 6 7 10]
}

func Example_sortFloat16() {
	ds := Array[float32]{4.7, 1.0, 2.6, 5.8, 7.1, 3.1, -5.5}
	result := Sort(ds)
	fmt.Println(result)
	// Output:
	// [-5.5 1 2.6 3.1 4.7 5.8 7.1]
}

func BenchmarkSort(b *testing.B) {

	benchcases := []struct {
		input Array[int]
	}{
		{input: Array[int]{400, 2, 11, 55}},
		{input: Array[int]{400, 2, 11, 55, 3, -50, 7}},
		{input: Array[int]{400, 2, 11, 55, 3, -50, 7, -100, 40, 9}},
		{input: Array[int]{400, 2, 11, 55, 3, -50, 7, -100, 40, 9, 70, 20, -70}},
		{input: Array[int]{400, 2, 11, 55, 3, -50, 7, -100, 40, 9, 70, 20, -70, 66, 90, 8}},
	}

	for _, bench := range benchcases {
		b.Run(fmt.Sprintf("input_%d", len(bench.input)), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				Sort(bench.input)
			}
		})
	}
}
