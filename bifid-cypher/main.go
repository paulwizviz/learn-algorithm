package main

import (
	"fmt"
	"log"
	"os"
)

type Cypher struct {
	rowIndex int
	colIndex int
	char     rune
}

func cyphers() []Cypher {
	return []Cypher{
		{rowIndex: 1, colIndex: 1, char: 'B'},
		{rowIndex: 1, colIndex: 2, char: 'G'},
		{rowIndex: 1, colIndex: 3, char: 'W'},
		{rowIndex: 1, colIndex: 4, char: 'K'},
		{rowIndex: 1, colIndex: 5, char: 'Z'},
		{rowIndex: 2, colIndex: 1, char: 'Q'},
		{rowIndex: 2, colIndex: 2, char: 'P'},
		{rowIndex: 2, colIndex: 3, char: 'N'},
		{rowIndex: 2, colIndex: 4, char: 'D'},
		{rowIndex: 2, colIndex: 5, char: 'S'},
		{rowIndex: 3, colIndex: 1, char: 'I'},
		{rowIndex: 3, colIndex: 2, char: 'O'},
		{rowIndex: 3, colIndex: 3, char: 'A'},
		{rowIndex: 3, colIndex: 4, char: 'X'},
		{rowIndex: 3, colIndex: 5, char: 'E'},
		{rowIndex: 4, colIndex: 1, char: 'F'},
		{rowIndex: 4, colIndex: 2, char: 'C'},
		{rowIndex: 4, colIndex: 3, char: 'L'},
		{rowIndex: 4, colIndex: 4, char: 'U'},
		{rowIndex: 4, colIndex: 5, char: 'M'},
		{rowIndex: 5, colIndex: 1, char: 'T'},
		{rowIndex: 5, colIndex: 2, char: 'H'},
		{rowIndex: 5, colIndex: 3, char: 'Y'},
		{rowIndex: 5, colIndex: 4, char: 'V'},
		{rowIndex: 5, colIndex: 5, char: 'R'},
	}
}

func charToCypher(input rune, cyphers []Cypher) (Cypher, error) {

	for _, seed := range cyphers {
		if seed.char == input {
			return seed, nil
		}
	}

	return Cypher{rowIndex: 0, colIndex: 0, char: ' '}, fmt.Errorf("Input character not part of seeds")
}

func posToCypher(rowNo int, colNo int, cyphers []Cypher) (Cypher, error) {

	for _, seed := range cyphers {
		if (seed.rowIndex == rowNo) && (seed.colIndex == colNo) {
			return seed, nil
		}
	}

	return Cypher{rowIndex: 0, colIndex: 0, char: ' '}, fmt.Errorf("Input rows and columns not part of seeds")
}

func planTextToCypherCoordinate(input string, cyphers []Cypher) ([]Cypher, error) {

	coordinates := []Cypher{}
	for _, ch := range input {
		coordinate, err := charToCypher(ch, cyphers)
		if err != nil {
			return []Cypher{}, err
		}
		coordinates = append(coordinates, coordinate)
	}
	return coordinates, nil
}

func plainToCypherText(input string, cyphers []Cypher) (string, error) {

	inputCoords, err := planTextToCypherCoordinate(input, cyphers)
	if err != nil {
		return "", err
	}

	flattenRow := []Cypher{}
	flattenCol := []Cypher{}
	for index := 0; index < len(inputCoords); index++ {
		currRow := inputCoords[index].rowIndex
		currCol := inputCoords[index].colIndex
		index++
		nextRow := inputCoords[index].rowIndex
		nextCol := inputCoords[index].colIndex
		coordinateRow, err := posToCypher(currRow, nextRow, cyphers)
		if err != nil {
			return "", err
		}
		flattenRow = append(flattenRow, coordinateRow)

		coordinateCol, err := posToCypher(currCol, nextCol, cyphers)
		if err != nil {
			return "", err
		}
		flattenCol = append(flattenCol, coordinateCol)
	}

	merged := []Cypher{}
	merged = append(merged, flattenRow...)
	merged = append(merged, flattenCol...)

	cypherText := []rune{}
	for _, item := range merged {
		cypherText = append(cypherText, item.char)
	}

	return string(cypherText), nil
}

func cypherToCypherCoordinate(input string, cyphers []Cypher) ([]Cypher, error) {

	coordinates := []Cypher{}
	for _, ch := range input {
		coordinate, err := charToCypher(ch, cyphers)
		if err != nil {
			return []Cypher{}, err
		}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func cypherToPlainText(input string, cyphers []Cypher) (string, error) {
	
	cypherCoords, err := cypherToCypherCoordinate(input, cyphers)
	if err != nil {
		return "", err
	}

	startIndexForCol := (len(cypherCoords) - len(cypherCoords)%2)/2 + len(cypherCoords)%2
	rows := []int{}
	cols := []int{}
	if len(cypherCoords)%2 != 0 {
		for index := 0; index < startIndexForCol; index++{
			if index == startIndexForCol-1{
				rows = append(rows, cypherCoords[index].rowIndex)
				cols = append(cols, cypherCoords[index].colIndex)
			}else {
				rows = append(rows, cypherCoords[index].rowIndex)
				rows = append(rows, cypherCoords[index].colIndex)
			}
		}
		for index := startIndexForCol;index < len(cypherCoords); index++{
			cols = append(cols, cypherCoords[index].rowIndex)
			cols = append(cols, cypherCoords[index].colIndex)
		}
	}else{
		for index := 0; index < len(cypherCoords); index++{
			rows = append(rows, cypherCoords[index].rowIndex)
		}
	}

	coordinates := []Cypher{}
	for index := 0; index < len(rows); index++{
		coordinate, err := posToCypher(rows[index], cols[index], cyphers)
		if err != nil{
			return "", err
		}
		coordinates = append(coordinates, coordinate)
	}

	plainText := []rune{}
	for _, coordinate := range coordinates{
		plainText = append(plainText, coordinate.char)
	}

	return string(plainText), nil
}

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Minimum of two arguments")
	}

}
