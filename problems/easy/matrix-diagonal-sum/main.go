package main

import "fmt"

func main() {
	fmt.Println(diagonalSum([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}))
}

/*

1 2 3    mat[row0][0]    1      mat[row0][2]     3
4 5 6	 mat[row1][1]    5      mat[row1][1]     5
7 8 9	 mat[row2][2]    9      mat[row2][0]     7

*/

func diagonalSum(mat [][]int) int {
	var sum int
	for i := 0; i < len(mat); i++ {
		if i == len(mat)-1-i {
			sum += mat[i][i]
			continue
		}
		sum += mat[i][i] + mat[i][len(mat)-1-i]
	}

	return sum
}
