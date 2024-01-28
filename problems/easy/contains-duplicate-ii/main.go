package main

import "fmt"

func main() {
	r := containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
	fmt.Println(r)
}

/*

Given an integer array nums and an integer k,
 return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

*/

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func containsNearbyDuplicate(nums []int, k int) bool {
	store := make(map[int]int)
	for right := 0; right < len(nums); right++ {
		if _, ok := store[nums[right]]; ok {
			diff := abs((store[nums[right]] - right))
			if diff <= k {
				return true
			}
		}

		store[nums[right]] = right
	}

	return false
}
