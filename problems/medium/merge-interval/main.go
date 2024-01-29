package main

import (
	"fmt"
	"sort"
)

/*
Given an array of intervals where intervals[i] = [starti, endi],
merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

*/

func main() {
	r := merge([][]int{{11, 12}, {2, 3}, {5, 7}, {1, 4}})
	fmt.Println(r)
}

/*

1-4, 2-3, 5-7, 11-12
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	var out [][]int
	out = append(out, intervals[0])

	for i := 0; i < len(intervals); i++ {
		lastInterval := out[len(out)-1]

		if lastInterval[1] >= intervals[i][0] {
			lastInterval[1] = max(lastInterval[1], intervals[i][1])
		} else {
			out = append(out, intervals[i])
		}
	}

	return out
}
