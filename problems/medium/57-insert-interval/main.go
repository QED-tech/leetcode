package main

import "fmt"

func main() {

	fmt.Println(
		insert(
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{3, 8},
		),
	)

}

func insert(intervals [][]int, newInterval []int) [][]int {
	var (
		out [][]int
		i   int
		n   = len(intervals)
	)

	newStart, newEnd := newInterval[0], newInterval[1]

	for i < n && intervals[i][1] < newStart {
		out = append(out, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newEnd {
		currentStart, currentEnd := intervals[i][0], intervals[i][1]

		if currentStart < newStart {
			newStart = currentStart
		}

		if currentEnd > newEnd {
			newEnd = currentEnd
		}

		i++
	}

	out = append(out, []int{newStart, newEnd})

	for i < n {
		out = append(out, intervals[i])
		i++
	}

	return out
}
