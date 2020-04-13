package main

import "testing"

func TestSearchRootWithCorrectValue(t *testing.T) {
	root := NewTreeSingleRoot()
	var result bool
	Search(0, root, &result)
	if result != true {
		t.Error("Error not expected")
	}
}

func TestSearchRootWithInCorrectValue(t *testing.T) {
	root := NewTreeSingleRoot()
	var result bool
	Search(1, root, &result)
	if result != false {
		t.Error("Error not expected")
	}
}

func TestSearchLevel1WithCorrectValue(t *testing.T) {
	root := New1LevelTree()
	var result bool
	Search(1, root, &result)
	if result != true {
		t.Error("Error not expected")
	}

	Search(2, root, &result)
	if result != true {
		t.Error("Error not expected")
	}
}

func TestSearchLevel1WithInCorrectValue(t *testing.T) {
	root := New1LevelTree()
	var result bool
	Search(11, root, &result)
	if result != false {
		t.Error("Error not expected")
	}
}
