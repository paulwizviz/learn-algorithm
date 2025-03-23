package bigo

import (
	"fmt"
	"testing"
)

// mapData types and operations to be benchmarked
var mapData = func() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 10_000; i++ {
		m[i] = i
	}
	return m
}()

func mapItemExist(item int) bool {
	_, ok := mapData[item]
	return ok
}

// slice data types and operations to be benchmarked
var sliceData = func() []int {
	var slice []int
	for value := 0; value < 10_000; value++ {
		slice = append(slice, value)
	}
	return slice
}()

func sliceItemExist(item int) bool {
	for _, value := range sliceData {
		if value == item {
			return true
		}
	}
	return false
}

// Benchmarked scenarios for O(1) and O(N)
var gettests = []struct {
	input int
}{
	{input: 1},
	{input: 20},
	{input: 50},
	{input: 100},
	{input: 900},
	{input: 1050},
	{input: 5000},
	{input: 9999},
}

// O(1)
func BenchmarkO1(b *testing.B) {
	for _, test := range gettests {
		b.Run(fmt.Sprintf("input_%d", test.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				mapItemExist(test.input)
			}
		})
	}
}

// O(N)
func BenchmarkON(b *testing.B) {
	for _, test := range gettests {
		b.Run(fmt.Sprintf("input_%d", test.input), func(b *testing.B) {
			for count := 0; count < b.N; count++ {
				sliceItemExist(test.input)
			}
		})
	}
}
