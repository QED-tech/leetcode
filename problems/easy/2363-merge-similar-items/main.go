package main

import (
	"fmt"
	"sort"
)

func main() {

	items1 := [][]int{
		{1, 1},
		{4, 5},
		{3, 8},
	}

	items2 := [][]int{
		{3, 1},
		{1, 5},
	}

	fmt.Println(mergeSimilarItems(items1, items2))
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var out [][]int

	m := make(map[int]int)

	for _, pair := range items1 {
		value, weith := pair[0], pair[1]
		m[value] += weith
	}

	for _, pair := range items2 {
		value, weith := pair[0], pair[1]
		m[value] += weith
	}

	for key, value := range m {
		out = append(out, []int{key, value})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i][0] < out[j][0]
	})

	return out
}
