package main

import "fmt"

func main() {
	r := singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Println(r)
}

func singleNumber(nums []int) int {
	var single int

	for i := 0; i < len(nums); i++ {
		single ^= nums[i]
	}

	return single
}
