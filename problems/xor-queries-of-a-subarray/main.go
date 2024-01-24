package main

import "fmt"

/*
Input: arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
Output: [2,7,14,8]
Explanation:
The binary representation of the elements in the array are:
1 = 0001
3 = 0011
4 = 0100
8 = 1000
The XOR values for queries are:
[0,1] = 1 xor 3 = 2
[1,2] = 3 xor 4 = 7
[0,3] = 1 xor 3 xor 4 xor 8 = 14
[3,3] = 8
*/

func main() {
	r := xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}})
	fmt.Println(r)
}

func xorQueries(arr []int, queries [][]int) []int {
	var out []int

	for _, q := range queries {
		start, end := q[0], q[1]
		var acc int

		for end >= start {
			acc ^= arr[start]
			start++
		}

		out = append(out, acc)
	}

	return out
}
