package table

import "fmt"

// numericType is a generic type constrains consisting of Go numerics.
type numericType interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

// Matrix is a table
type Matrix[N numericType] interface {
	Set(xPos, yPos uint16, value N) error
	ValueAt(xPos, yPos uint16) (N, error)
	Rows() uint16
	Cols() uint16
}

// defaultMatrix is a variable-size matrix
type defaultMatrix[N numericType] struct {
	elems [][]N
	rows  uint16
	cols  uint16
}

func (m *defaultMatrix[N]) Set(xPos, yPos uint16, value N) error {

	if m.rows-1 < xPos {
		return fmt.Errorf("Position greater than max rows")
	}

	if m.cols-1 < yPos {
		return fmt.Errorf("Position greater than max columns")
	}

	m.elems[xPos][yPos] = value
	return nil
}

func (m defaultMatrix[N]) ValueAt(xPos, yPos uint16) (N, error) {

	if m.rows-1 < xPos {
		return 0, fmt.Errorf("Position greater than max rows")
	}

	if m.cols-1 < yPos {
		return 0, fmt.Errorf("Position greater than max columns")
	}

	return m.elems[xPos][yPos], nil
}

func (m defaultMatrix[N]) Rows() uint16 {
	return m.cols
}

func (m defaultMatrix[N]) Cols() uint16 {
	return m.rows
}

func NewMatrix[N numericType](rows, cols uint16) Matrix[N] {
	elems := make([][]N, rows)
	for xPos := 0; xPos < len(elems); xPos++ {
		elems[xPos] = make([]N, cols)
	}
	return &defaultMatrix[N]{
		elems: elems,
		rows:  rows,
		cols:  cols,
	}
}
