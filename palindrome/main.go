package main

import (
	"log"
)

var data1 = []string{"abcdef", "anna", "cijkl", "civic"}
var data2 = []string{"abcded", "aann", "cijkl", "ciivc"}

func isPalindrome(input string) bool {

	var lenOfCharset = len(input)
	forwardCharset := []rune(input)
	tempCharset := []rune(input)

	var reverseCharset []rune
	reverseCharset = make([]rune, lenOfCharset, lenOfCharset)

	for i := 0; i < lenOfCharset; i++ {
		reverseCharset[i] = tempCharset[lenOfCharset-1-i]
	}

	for i := 0; i < lenOfCharset; i++ {
		if forwardCharset[i] != reverseCharset[i] {
			return false
		}
	}

	return true
}

func main() {
	for _, word := range data1 {
		log.Printf("Word: %v is palindrome: %v", word, isPalindrome(word))
	}
}
