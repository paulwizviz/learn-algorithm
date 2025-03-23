package bigo

import (
	"fmt"
	"testing"
)

// operations to determine if two slices are duplicated
func isDupSlice(d1, d2 []int) bool {
	for oIndex, outer := range d1 {
		for iIndex, inner := range d2 {
			if oIndex == iIndex {
				if outer != inner {
					return false
				}
			}
		}
	}
	return true
}

// O(N^2)
var dupTest = []struct {
	d1 []int
	d2 []int
	tc int
}{
	{
		d1: []int{1, 2},
		d2: []int{1, 2},
		tc: 1,
	},
	{
		d1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		d2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		tc: 2,
	},
	{
		d1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		d2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		tc: 3,
	},
}

func BenchmarkN2(b *testing.B) {
	for _, test := range dupTest {
		b.Run(fmt.Sprintf("input_%d", test.tc), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				isDupSlice(test.d1, test.d2)
			}
		})
	}
}
