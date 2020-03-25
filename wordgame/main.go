package main

import (
	"bufio"
	"log"
	"os"
)

type WordGame struct {
	dictionary []string
	scores     []map[string]int
}

func (wg *WordGame) WordEntry(word string) {
	wg.scores = findProperWordsAndScore(word, wg.dictionary)
	log.Println(wg.scores)
}

func (wg *WordGame) WordAtPosition(position int) string {
	if position < len(wg.scores)-1 {
		score := wg.scores[position]
		for w, _ := range score {
			log.Println(w)
			return w
		}
	}
	return ""
}

func (wg *WordGame) ScoreAtPosition(position int) int {
	return 0
}

func newWordGame(path string) *WordGame {
	dictionary, err := scanWords(path)
	if err != nil {
		log.Fatalf("Failed to create dictionary. %v", err)
	}

	return &WordGame{
		dictionary: dictionary,
		scores:     []map[string]int{},
	}
}

func scanWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}

func isValidWord(input string, validWords []string) bool {

	for _, vw := range validWords {
		if input == vw {
			return true
		}
	}

	return false
}

func swapCharacters(input []rune, currentIndex int, nextIndex int) string {
	var tmpChar rune
	tmpChar = input[currentIndex]
	input[currentIndex] = input[nextIndex]
	input[nextIndex] = tmpChar
	return string(input)
}

func shiftFirstCharacterToFormWords(input string) []string {
	sliceOfWords := []string{}
	inputRuneSlice := []rune(input)

	locationOfLastCharacter := len(input) - 1
	locationOfFirstCharacter := 0
	for _, _ = range inputRuneSlice {
		locationOfNextCharacter := locationOfFirstCharacter + 1
		if locationOfNextCharacter <= locationOfLastCharacter {
			word := swapCharacters(inputRuneSlice, locationOfFirstCharacter, locationOfNextCharacter)
			sliceOfWords = append(sliceOfWords, word)
			locationOfFirstCharacter++
		}
	}

	return sliceOfWords
}

func rearrangeWords(input string) []string {

	numberOfRunes := len(input)
	wordCombinations := []string{}

	for counter := 1; counter <= numberOfRunes; counter++ {
		if counter == 1 {
			wordCombinations = append(wordCombinations, shiftFirstCharacterToFormWords(input)...)
		} else {
			wordCombinations = append(wordCombinations, shiftFirstCharacterToFormWords(wordCombinations[len(wordCombinations)-1])...)
		}
	}

	return wordCombinations
}

func findProperWordsAndScore(input string, dictionary []string) []map[string]int {

	scores := []map[string]int{}
	for _, word := range rearrangeWords(input) {
		log.Println(word)
		score := map[string]int{}
		if isValidWord(word, dictionary) {
			score[word] = len(word)
			scores = append(scores, score)
		}
	}
	return scores
}

func main() {
	wg := newWordGame("./words.txt")
	wg.WordEntry("abracadbra")
	log.Println(wg.WordAtPosition(1))
}
