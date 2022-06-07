package main

import (
	"flag"
	"fmt"
)

type inttypes interface {
	uint8 | uint | int
}

type Board[d inttypes] struct {
	grids [][]d
	rows  uint
	cols  uint
}

func (b *Board[d]) Set(xPos, yPos uint, value d) error {

	if b.rows-1 < xPos {
		return fmt.Errorf("Position greater than max rows")
	}

	if b.cols-1 < yPos {
		return fmt.Errorf("Position greater than max columns")
	}

	b.grids[xPos][yPos] = value
	return nil
}

func (b Board[d]) ValueAt(xPos, yPos uint) (d, error) {

	if b.rows-1 < xPos {
		return 0, fmt.Errorf("Position greater than max rows")
	}

	if b.cols-1 < yPos {
		return 0, fmt.Errorf("Position greater than max columns")
	}

	return b.grids[xPos][yPos], nil
}

func NewBoard[d inttypes](rows, cols uint) *Board[d] {

	grids := make([][]d, rows)
	for xPos := 0; xPos < len(grids); xPos++ {
		grids[xPos] = make([]d, cols)
	}

	return &Board[d]{
		grids: grids,
		rows:  rows,
		cols:  cols,
	}
}

func hit[d inttypes](b *Board[d], xPos, yPos uint) (bool, error) {

	result, err := b.ValueAt(xPos, yPos)
	if err != nil {
		return false, fmt.Errorf("Out of bound")
	}

	if result != 1 {
		return false, nil
	}

	return true, nil
}

func main() {
	board := NewBoard[uint8](3, 5)

	board.Set(0, 0, 0)
	board.Set(0, 1, 1)
	board.Set(0, 2, 0)
	board.Set(0, 3, 0)
	board.Set(0, 4, 0)
	board.Set(1, 0, 0)
	board.Set(1, 1, 1)
	board.Set(1, 2, 0)
	board.Set(1, 3, 0)
	board.Set(1, 4, 0)
	board.Set(2, 0, 0)
	board.Set(2, 1, 0)
	board.Set(2, 2, 0)
	board.Set(2, 3, 0)
	board.Set(2, 4, 0)

	xPosPtr := flag.Uint("xpos", 0, "x position")
	yPosPtr := flag.Uint("ypos", 0, "y position")

	flag.Parse()

	result, _ := hit(board, *xPosPtr, *yPosPtr)
	fmt.Println(result)

}
