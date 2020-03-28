package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCharToCypher(t *testing.T) {
	input := 'A'
	expected := Cypher{
		rowIndex: 3,
		colIndex: 3,
		char:     'A',
	}
	actual, err := charToCypher(input, cyphers())
	if err != nil {
		t.Fatalf("Unable to find base character co-ordinates %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v Got: %v", expected, actual)
	}
}

func TestPosToCypher(t *testing.T) {
	colNo := 3
	rowNo := 3
	expected := Cypher{
		rowIndex: 3,
		colIndex: 3,
		char:     'A',
	}
	actual, err := posToCypher(rowNo, colNo, cyphers())
	if err != nil {
		t.Fatalf("Unable to find base character co-ordinates %v", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v Got: %v", expected, actual)
	}
}

func TestTextToCypherCoordinate(t *testing.T) {
	dataSet := []struct {
		input    string
		expected []Cypher
	}{
		{
			input: "MARYLAND",
			expected: []Cypher{
				{
					rowIndex: 4,
					colIndex: 5,
					char:     'M',
				},
				{
					rowIndex: 3,
					colIndex: 3,
					char:     'A',
				},
				{
					rowIndex: 5,
					colIndex: 5,
					char:     'R',
				},
				{
					rowIndex: 5,
					colIndex: 3,
					char:     'Y',
				},
				{
					rowIndex: 4,
					colIndex: 3,
					char:     'L',
				},
				{
					rowIndex: 3,
					colIndex: 3,
					char:     'A',
				},
				{
					rowIndex: 2,
					colIndex: 3,
					char:     'N',
				},
				{
					rowIndex: 2,
					colIndex: 4,
					char:     'D',
				},
			},
		},
	}
	for _, data := range dataSet {
		actual, err := planTextToCypherCoordinate(data.input, cyphers())
		if err != nil {
			t.Fatalf("Unable to produce coordinate: %v", err)
		}

		if !reflect.DeepEqual(data.expected, actual) {
			t.Errorf("Expected: %v Got: %v", data.expected, actual)
		}
	}
}

func TestPlainToCypherText(t *testing.T) {
	dataSet := []struct {
		input    string
		expected string
	}{
		{
			input:    "MARYLAND",
			expected: "LRLPYYAX",
		},
	}

	for _, data := range dataSet {
		actual, err := plainToCypherText(data.input, cyphers())
		if err != nil {
			t.Fatalf("Unable to generate cypher. %v", err)
		}

		if strings.Compare(data.expected, actual) != 0 {
			t.Logf("Expected: %v Got: %v", data.expected, actual)
		}
	}

}

func TestCypherToCypherCoordinate(t *testing.T){
	dataSet := []struct{
		input string
		expected []Cypher
	}{
		{
			input: "DXETE",
			expected: []Cypher{
				Cypher{
					rowIndex: 2,
					colIndex: 4,
					char: 'D',
				},
				Cypher{
					rowIndex: 3,
					colIndex: 4,
					char: 'X',
				},
				Cypher{
					rowIndex: 3,
					colIndex: 5,
					char: 'E',
				},
				Cypher{
					rowIndex: 5,
					colIndex: 1,
					char: 'T',
				},
				Cypher{
					rowIndex: 3,
					colIndex: 5,
					char: 'E',
				},
			},
		},
	}

	for _, data := range dataSet{
		actual, err := cypherToCypherCoordinate(data.input, cyphers())
		if err != nil{
			t.Fatalf("Unable to convert cypher text to coordinate. %v", err)
		}
		if !reflect.DeepEqual(data.expected, actual){
			t.Errorf("Expected: %v Got: %v", data.expected, actual)
		}
	}

}

func TestCypherToPlainText(t *testing.T) {
	dataSet := []struct{
		input string
		expected string
	}{
		{
			input: "DXETE",
			expected: "SMILE",
		},
	}

	for _, data := range dataSet{
		actual, err := cypherToPlainText(data.input, cyphers())
		if err != nil {
			t.Fatalf("Unable to generate plain text. %v", err)
		}

		if strings.Compare(data.expected, actual) != 0 {
			t.Logf("Expected: %v Got: %v", data.expected, actual)
		}
	}
}
