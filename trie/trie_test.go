package trie

import (
	"fmt"
	"strconv"
)

func Example_isLeaf() {
	node := NewDefaultNode()
	fmt.Printf("Is a terminated edge? %v", node.InsertRune('a').IsLeaf())

	// Output:
	// Is a terminated edge? true
}

func Example_matchedRune() {
	trie := NewDefaultNode()
	node := trie.InsertRune('a')
	node.InsertRune('b')

	isEnd, isMatched := trie.MatchedRune('a')
	fmt.Printf("Matching first rune. isEnd: %v isMatched: %v\n", isEnd, isMatched)

	isEnd, isMatched = trie.MatchedRune('c')
	fmt.Printf("Not matching first rune. isEnd: %v isMatched: %v\n", isEnd, isMatched)

	isEnd, isMatched = node.MatchedRune('b')
	fmt.Printf("Matching terminated rune. isEnd: %v isMatched: %v\n", isEnd, isMatched)

	isEnd, isMatched = node.MatchedRune('c')
	fmt.Printf("Not matching terminated rune. isEnd: %v isMatched: %v", isEnd, isMatched)

	// Output:
	// Matching first rune. isEnd: false isMatched: true
	// Not matching first rune. isEnd: false isMatched: false
	// Matching terminated rune. isEnd: true isMatched: true
	// Not matching terminated rune. isEnd: false isMatched: false
}

func Example_populateTrie() {

	trie := NewDefaultNode()
	Populate(trie, "abc")

	var node Node = trie
	var res []string
	for {
		runes := node.Runes()
		str := strconv.QuoteRune(runes[0])
		res = append(res, str)
		node = node.NextNode(runes[0])
		if node.IsLeaf() {
			break
		}
	}

	fmt.Println(res)

	// Output:
	// ['a' 'b' 'c']
}

func Example_searchTrie() {

	trie := NewDefaultNode()
	Populate(trie, "ant")
	Populate(trie, "another")

	result := VerifyWord(trie, "ant")
	fmt.Printf("Verify that 'ant' is found in this trie and returns error. %v\n", result)

	result = VerifyWord(trie, "another")
	fmt.Printf("Verify that 'another' is found in this trie and returns error. %v\n", result)

	result = VerifyWord(trie, "an")
	fmt.Println(result)

	// Output:
	// Verify that 'ant' is found in this trie and returns error. <nil>
	// Verify that 'another' is found in this trie and returns error. <nil>
	// Trie search error. Unable to match criteria: an
}

func Example_deleteTrie() {
	trie := NewDefaultNode()
	Populate(trie, "ant")
	Populate(trie, "another")

	result := DeleteWord(trie, "an")
	fmt.Println(result)

	result = DeleteWord(trie, "another")
	fmt.Printf("Word 'another' is deleted from trie. Error: %v\n", result)

	result = VerifyWord(trie, "another")
	fmt.Printf("Veification of word 'another' yield no Error. %v", result)

	// Output:
	// Trie search error. Unable to delete word: an
	// Word 'another' is deleted from trie. Error: <nil>
	// Veification of word 'another' yield no Error. <nil>

}
