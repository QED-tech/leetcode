package main

import (
	"fmt"
)

var validElements = map[byte]struct{}{
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
}

func isValidSudoku(board [][]byte) bool {

	var (
		empty   byte = '.'
		rowsLen      = len(board)
		colsLen      = len(board[0])
	)

	rows := [9][9]byte{}
	cols := [9][9]byte{}
	boxes := [9][9]byte{}

	for row := 0; row < rowsLen; row++ {
		for col := 0; col < colsLen; col++ {
			el := board[row][col]
			if el == empty {
				continue
			}

			if _, ok := validElements[el]; !ok {
				return false
			}

			rows[row][col] = el
			cols[col][row] = el

			box := (row/3)*3 + col/3
			indexInBox := (row%3)*3 + (col % 3)

			boxes[box][indexInBox] = el
		}
	}

	isValidSector := func(sector [9][9]byte) bool {
		for _, s := range sector {
			m := make(map[byte]struct{})

			for _, el := range s {
				if el == 0 {
					continue
				}

				if _, ok := m[el]; ok {
					return false
				}

				m[el] = struct{}{}
			}

		}

		return true
	}

	if !isValidSector(rows) {
		return false
	}
	if !isValidSector(cols) {
		return false
	}
	if !isValidSector(boxes) {
		return false
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},

		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},

		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	result := isValidSudoku(board)
	fmt.Println(result)
}
