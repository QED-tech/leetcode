package main

import "fmt"

func main() {
	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}

	result := countNegatives(grid)
	fmt.Println(result)
}

func countNegatives(grid [][]int) int {
	var count int

	rows := len(grid)
	cols := len(grid[0])

	row := 0
	col := cols - 1

	for row < rows && col >= 0 {
		if grid[row][col] < 0 {
			count += rows - row

			col--
		} else {
			row++
		}
	}

	return count
}
