package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	low, hight := 0, len(nums)-1

	for low <= hight {
		mid := (hight + low) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
