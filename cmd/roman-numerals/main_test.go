package main

import (
	"reflect"
	"testing"
)

func TestSymbolicValues(t *testing.T) {

	input := "IVXLCDM"
	expected := []int{1, 5, 10, 50, 100, 500, 1000}
	actual := romanToSymbolicValues(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v Actual: %v", expected, actual)
	}

}

type keyvalue struct {
	key   string
	value int
}

func TestConvertNumeralsToInteger(t *testing.T) {

	testData := []keyvalue{
		{key: "MCMLXXVI", value: 1976},
		{key: "MMXVIII", value: 2018},
		{key: "MCMLXV", value: 1965},
	}

	for _, data := range testData {

		actual := convertSymbolsToNumbericValues(data.key)
		expected := data.value
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("Expected: %v Got: %v", expected, actual)
		}
	}
}
