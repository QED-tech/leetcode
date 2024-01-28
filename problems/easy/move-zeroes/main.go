package main

import "fmt"

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

*/

func main() {
	l := []int{1, 2, 0, 0, 0, 0, 3, 4}
	moveZeroes(l)
	fmt.Println(l)
}

func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			left = i
			for left < len(nums)-1 && nums[left] == 0 {
				left++
			}
			nums[left], nums[i] = nums[i], nums[left]
		}
	}
}
