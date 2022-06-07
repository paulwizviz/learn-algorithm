package main

import "fmt"

func Example_newBoard() {

	board := NewBoard[uint8](2, 3)
	fmt.Println(board)

	// Output:
	// &{[[0 0 0] [0 0 0]] 2 3}
}

func Example_populateBoard() {
	board := NewBoard[uint8](2, 3)

	board.Set(0, 0, 1)
	board.Set(0, 1, 2)
	board.Set(0, 2, 3)
	board.Set(1, 0, 4)
	board.Set(1, 1, 5)
	board.Set(1, 2, 6)

	fmt.Println(board)

	err := board.Set(2, 0, 10)
	fmt.Println(err)

	err = board.Set(0, 3, 11)
	fmt.Println(err)

	// Output:
	// &{[[1 2 3] [4 5 6]] 2 3}
	// Position greater than max rows
	// Position greater than max columns
}

func Example_valueAt() {
	board := NewBoard[uint8](2, 2)

	board.Set(0, 0, 1)
	board.Set(0, 1, 2)
	board.Set(1, 0, 3)
	board.Set(1, 1, 4)

	result, _ := board.ValueAt(1, 0)
	fmt.Println(result)

	_, err := board.ValueAt(2, 0)
	fmt.Println(err)

	_, err = board.ValueAt(0, 2)
	fmt.Println(err)

	// Output:
	// 3
	// Position greater than max rows
	// Position greater than max columns
}
