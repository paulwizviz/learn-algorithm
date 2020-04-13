package main

type node struct {
	value     int
	rightNode *node
	leftNode  *node
}

func AddLeafNodes(root *node, left *node, right *node) {
	if left != nil {
		root.leftNode = left
	}
	if right != nil {
		root.rightNode = right
	}
}

func NewTreeSingleRoot() *node {
	return &node{
		value:     0,
		leftNode:  nil,
		rightNode: nil,
	}
}

func New1LevelTree() *node {
	root := node{value: 0, leftNode: nil, rightNode: nil}
	left := node{value: 1, leftNode: nil, rightNode: nil}
	right := node{value: 2, leftNode: nil, rightNode: nil}
	AddLeafNodes(&root, &left, &right)
	return &root
}

func New2LevelTree() *node {
	root := node{value: 0, leftNode: nil, rightNode: nil}
	left := node{value: 1, leftNode: nil, rightNode: nil}
	right := node{value: 2, leftNode: nil, rightNode: nil}
	AddLeafNodes(&root, &left, &right)
	left11 := node{value: 11, leftNode: nil, rightNode: nil}
	right12 := node{value: 12, leftNode: nil, rightNode: nil}
	AddLeafNodes(&left, &left11, &right12)
	left21 := node{value: 21, leftNode: nil, rightNode: nil}
	right22 := node{value: 22, leftNode: nil, rightNode: nil}
	AddLeafNodes(&right, &left21, &right22)
	return &root
}

func Search(expectedValue int, root *node, result *bool) {

	left := root.leftNode
	right := root.rightNode
	if root.value != expectedValue {
		if left != nil {
			Search(expectedValue, left, result)
			if *result == true {
				return
			}
		}
		if right != nil {
			Search(expectedValue, right, result)
			if *result == true {
				return
			}
		}
		*result = false
		return
	}
	*result = true
}
