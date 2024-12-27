// A simple routine to form words from rearranging seed characters

package main

import (
	"log"
)

func rearrangeRune(input []rune) []string {

	wordCombinations := []string{}

	inputClone := make([]rune, len(input))
	copy(inputClone, input)

	indexOfFirstCharater := 0
	for counter := 0; counter < len(input)-1; counter++ {

		if indexOfFirstCharater+1 < len(input) {
			tmp := inputClone[indexOfFirstCharater]
			inputClone[indexOfFirstCharater] = inputClone[indexOfFirstCharater+1]
			inputClone[indexOfFirstCharater+1] = tmp
			indexOfFirstCharater++
		}

		wordCombinations = append(wordCombinations, string(inputClone))
	}
	return wordCombinations
}

func main() {

	data := "abc"

	wordCombinations := []string{data}
	wordCombinations = append(wordCombinations, rearrangeRune([]rune(wordCombinations[0]))...)
	log.Println(wordCombinations)

	for count := 1; count < len(data); count++ {
		wordCombinations = append(wordCombinations, rearrangeRune([]rune(wordCombinations[len(wordCombinations)-1]))...)
		log.Println(wordCombinations)
	}
}
