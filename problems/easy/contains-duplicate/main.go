package main

import "fmt"

func main() {
	r := containsDuplicate([]int{1, 2, 3, 1})
	fmt.Println(r)
}

func containsDuplicate(nums []int) bool {
	stor := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := stor[num]; ok {
			return true
		}
		stor[num] = struct{}{}
	}

	return false
}
