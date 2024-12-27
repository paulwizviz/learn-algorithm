// A classic FizzBuzz algorithm to demonstrate TDD techniques

package main

import (
	"log"
	"strconv"
)

func convertNumber(number int) string {

	if number%15 == 0 {
		return "FizzBuzz"
	}

	if number%3 == 0 {
		return "Fizz"
	}

	if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

func main() {
	for i := 1; i < 101; i++ {
		log.Printf("Entry: %d Output: %s\n", i, convertNumber(i))
	}
}
