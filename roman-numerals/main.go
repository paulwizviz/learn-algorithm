package main

import (
	"flag"
	"log"
	"strings"
)

func main() {

	roman := flag.String("roman", "", "Year in roman numerals")
	flag.Parse()

	if strings.Compare(*roman, "") == 0 {
		log.Fatal("Failed to supply a value")
	}

	log.Println(convertSymbolsToNumbericValues(*roman))

}

func romanToSymbolicValues(input string) []int {

	symbols := []rune(input)
	symbolicValues := []int{}

	for _, symbol := range symbols {
		switch symbol {
		case 'I':
			symbolicValues = append(symbolicValues, 1)
		case 'V':
			symbolicValues = append(symbolicValues, 5)
		case 'X':
			symbolicValues = append(symbolicValues, 10)
		case 'L':
			symbolicValues = append(symbolicValues, 50)
		case 'C':
			symbolicValues = append(symbolicValues, 100)
		case 'D':
			symbolicValues = append(symbolicValues, 500)
		case 'M':
			symbolicValues = append(symbolicValues, 1000)
		default:
			// Do nothing
		}
	}
	return symbolicValues
}

func convertSymbolsToNumbericValues(input string) int {

	symbolicValues := romanToSymbolicValues(input)

	accum := 0
	for index := 0; index <= len(symbolicValues)-1; index++ {
		currentValue := symbolicValues[index]
		var nextValue int
		if index+1 <= len(symbolicValues)-1 {
			nextValue = symbolicValues[index+1]
		}
		if currentValue >= nextValue {
			accum = accum + currentValue
		} else {
			compoundValue := nextValue - currentValue
			accum = accum + compoundValue
			index++
		}

	}
	return accum
}
