package trie

import "unicode"

// This is not serializable
type defaultNode struct {
	edges map[rune]Node
}

func (d *defaultNode) InsertRune(data rune) Node {
	lrune := unicode.ToLower(data)
	node, ok := d.edges[lrune]
	if !ok {
		d.edges[lrune] = NewDefaultNode()
		return d.edges[lrune]
	}
	return node
}

func (d *defaultNode) IsLeaf() bool {
	return len(d.edges) == 0
}

func (d *defaultNode) MatchedRune(data rune) (isEnd bool, isMatched bool) {
	isMatched, isEnd = false, false
	nextNode, ok := d.edges[data]
	if !ok {
		return isEnd, isMatched
	}
	isEnd, isMatched = nextNode.IsLeaf(), true
	return isEnd, isMatched
}

func (d *defaultNode) NextNode(ch rune) (node Node) {
	node, ok := d.edges[ch]
	if !ok {
		return nil
	}
	return node
}

func (d *defaultNode) Runes() []rune {
	var runes []rune

	for r, _ := range d.edges {
		runes = append(runes, r)
	}

	return runes
}

// NewDefaultNode instantiate a default Trie node.
func NewDefaultNode() Node {
	return &defaultNode{
		edges: make(map[rune]Node),
	}
}
