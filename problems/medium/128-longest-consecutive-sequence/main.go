package main

import "fmt"

func main() {

	fmt.Println(longestConsecutive([]int{
		0,
		3,
		7,
		2,
		5,
		8,
		4,
		6,
		0,
		1,
	})) // 9
}

/*
	0,
	3,
	7,
	2,
	5,
	8,
	4,
	6,
	0,
	1,
*/

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}

	var res int

	for num := range m {
		seq := 1

		if _, ok := m[num-1]; ok {
			continue
		}

		for {
			if _, ok := m[num+seq]; ok {
				seq++
				continue
			}

			res = max(res, seq)
			break
		}
	}

	return res
}
