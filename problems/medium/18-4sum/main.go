package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

/*
1,0,-1,0,-2,2


|i|	|z|	|l|			|r|
-2	-1	0	0	1	2
0	1	2	3	4	5

if i + l + r + z == target ? l++ r--



for i := 0; i < len(nums) - 3; i++ {
	left := i + 1
	right := len(nums) - 2

	for z := i + 1; z < len(nums)-2; z++ {

	left := z + 1
	right := n - 1

	for left < right {

	}
	}
}

*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	m := make(map[[4]int]struct{})

	var (
		out [][]int
		n   = len(nums)
	)

	for i := 0; i < n-3; i++ {
		for z := i + 1; z < n; z++ {
			left := z + 1
			right := n - 1
			for left < right {
				sum := nums[i] + nums[z] + nums[left] + nums[right]

				if sum == target {
					val := [4]int{nums[i], nums[z], nums[left], nums[right]}

					if _, ok := m[val]; !ok {
						out = append(out, []int{nums[i], nums[z], nums[left], nums[right]})
						m[val] = struct{}{}
					}

					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return out
}
