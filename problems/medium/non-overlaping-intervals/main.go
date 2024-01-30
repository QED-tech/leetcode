package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var counter int

	for i := 1; i < len(intervals); i++ {
		currentInterval := intervals[i]
		prevInterval := intervals[i-1]

		if prevInterval[1] > currentInterval[0] {
			intervals[i] = intervals[i-1]
			counter++
		}
	}

	return counter
}
