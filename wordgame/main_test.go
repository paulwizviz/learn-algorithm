package main

import (
	"reflect"
	"testing"
)

func setupDictionary(t *testing.T) []string {

	words, err := scanWords("./words.txt")
	if err != nil {
		t.Fatalf("Unable to setup dictionary: %v", err)
	}
	return words
}

func TestCheckWord(t *testing.T) {

	dictionary := setupDictionary(t)

	validWords := make([]string, 0)

	for _, input := range []string{"azbsd", "zoo", "accompanying"} {
		if isValidWord(input, dictionary) {
			validWords = append(validWords, input)
		}
	}

	if len(validWords) != 2 {
		t.Fatalf("Expected: %v Got: %v", 1, len(validWords))
	}
}

func TestSwapCharacters(t *testing.T) {
	input := []rune{'a', 'b'}

	output := swapCharacters(input, 0, 1)
	if output != "ba" {
		t.Fatalf("Expected: ba Got: %v", output)
	}
}

func TestShiftFirstCharacter(t *testing.T) {
	input := "abc"
	actualSlice := shiftFirstCharacterToFormWords(input)
	expectedSlice := []string{"bac", "bca"}
	if !reflect.DeepEqual(expectedSlice, actualSlice) {
		t.Fatalf("Expected: %v Got: %v", expectedSlice, actualSlice)
	}
}

func TestRearrangeWork(t *testing.T) {
	input := "abc"
	expectedSlice := []string{"bac", "bca", "cba", "cab", "acb", "abc"}
	actualSlice := rearrangeWords(input)
	if !reflect.DeepEqual(expectedSlice, actualSlice) {
		t.Fatalf("Expected: %v Got: %v", expectedSlice, actualSlice)
	}
}

func TestFindProperWord(t *testing.T) {
	input := "abc"
	dictionary := setupDictionary(t)
	expected := map[string]int{"cab": 3, "abc": 3}
	actual := findProperWordsAndScore(input, dictionary)
	if reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v, Got: %v", expected, actual)
	}
}
