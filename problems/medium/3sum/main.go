package main

import (
	"fmt"
	"sort"
)

func main() {
	r := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(r)
}

func threeSum(nums []int) [][]int {
	var out [][]int

	sort.Ints(nums)

	for left := 0; left < len(nums)-2; left++ {
		if left > 0 && nums[left-1] == nums[left] {
			continue
		}

		for right := len(nums) - 1; right > left; right-- {
			if right < len(nums)-1 && nums[right+1] == nums[right] {
				continue
			}

			needle := -(nums[right] + nums[left])

			if need := binarySearch(nums, needle, left+1, right-1); need != -1 {
				out = append(out, []int{nums[left], nums[right], needle})
			}

		}
	}

	return out
}

func binarySearch(list []int, needle int, low, hight int) int {
	for hight >= low {
		middle := (low + hight) / 2

		if list[middle] == needle {
			return middle
		}

		if needle > list[middle] {
			low = middle + 1
		} else {
			hight = middle - 1
		}
	}

	return -1
}
