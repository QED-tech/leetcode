package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{8, 9, 10, 11, -2, -1, 0, 1, 2, 3, 4}))
}

/*
               r
               l
           m
8, 9, 10, -2, -1, 0, 1, 2, 3, 4

if nums[m] > nums[r] | -1 > 4 ? 10 > -1 ?
 ...
else
 right = m

*/

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
