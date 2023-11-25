package table

import "fmt"

func Example_newMatrix() {
	vMatrix := NewMatrix[int16](2, 3)
	cols := vMatrix.Cols()
	rows := vMatrix.Rows()
	fmt.Printf("Rows: %v Cols: %v", rows, cols)

	// Output:
	// Rows: 3 Cols: 2
}

func Example_setMatrix() {
	vMatrix := NewMatrix[int16](2, 3)
	fmt.Println(vMatrix)

	err := vMatrix.Set(0, 0, 1)
	fmt.Println(vMatrix, err)

	err = vMatrix.Set(0, 1, 2)
	fmt.Println(vMatrix, err)

	err = vMatrix.Set(0, 2, 3)
	fmt.Println(vMatrix, err)

	err = vMatrix.Set(1, 0, 4)
	fmt.Println(vMatrix, err)

	err = vMatrix.Set(1, 1, 5)
	fmt.Println(vMatrix, err)

	err = vMatrix.Set(1, 2, 6)
	fmt.Println(vMatrix, err)

	err = vMatrix.Set(2, 0, 7)
	fmt.Println(vMatrix, err)

	err = vMatrix.Set(1, 3, 8)
	fmt.Println(vMatrix, err)

	// Output:
	// &{[[0 0 0] [0 0 0]] 2 3}
	// &{[[1 0 0] [0 0 0]] 2 3} <nil>
	// &{[[1 2 0] [0 0 0]] 2 3} <nil>
	// &{[[1 2 3] [0 0 0]] 2 3} <nil>
	// &{[[1 2 3] [4 0 0]] 2 3} <nil>
	// &{[[1 2 3] [4 5 0]] 2 3} <nil>
	// &{[[1 2 3] [4 5 6]] 2 3} <nil>
	// &{[[1 2 3] [4 5 6]] 2 3} Position greater than max rows
	// &{[[1 2 3] [4 5 6]] 2 3} Position greater than max columns
}

func Example_getMatrix() {
	vMatrix := NewMatrix[int16](2, 3)
	fmt.Printf("Initial matrix: %v\n", vMatrix)

	vMatrix.Set(0, 0, 1)
	value, _ := vMatrix.ValueAt(0, 0)
	fmt.Printf("Item for 0,0 is %v\n", value)
	value, _ = vMatrix.ValueAt(0, 1)
	fmt.Printf("Item for 0,1 is %v", value)

	// Output:
	// Initial matrix: &{[[0 0 0] [0 0 0]] 2 3}
	// Item for 0,0 is 1
	// Item for 0,1 is 0
}
