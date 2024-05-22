package main

import "testing"

func TestRecursionGenesis(t *testing.T) {
	var got = 0
	ReduceToZero(&got, 0)
	expected := 0
	if expected != got {
		t.Errorf("Expected: %d Got: %d", expected, got)
	}
}

func TestRecursionReduce10(t *testing.T) {
	var got = 0
	ReduceToZero(&got, 10)
	expected := 10
	if expected != got {
		t.Errorf("Expected: %d Got: %d", expected, got)
	}
}
