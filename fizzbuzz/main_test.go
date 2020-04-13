package main

import (
	"reflect"
	"testing"
)

func TestJustReturnANumber(t *testing.T) {
	expected := "2"
	got := convertNumber(2)

	if expected != got {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}

func TestDivisibleByThree(t *testing.T) {

	expected := "Fizz"
	got := convertNumber(3)

	if expected != got {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}

}

func TestDivisibleByFive(t *testing.T) {
	expected := "Buzz"
	got := convertNumber(5)

	if expected != got {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}

func TestDivisibleByFifteen(t *testing.T) {
	expected := "FizzBuzz"
	got := convertNumber(15)

	if expected != got {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}

func TestARangeOfData(t *testing.T) {

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz", "26", "Fizz", "28", "29", "FizzBuzz", "31", "32", "Fizz", "34", "Buzz", "Fizz"}

	got := []string{}
	for i := 1; i < 37; i++ {
		got = append(got, convertNumber(i))
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}
