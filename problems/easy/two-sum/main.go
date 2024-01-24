package main

import "fmt"

func main() {
	r := twoSum([]int{3, 2, 4}, 5)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	out, seen := make([]int, 2), make(map[int]int)

	for i, num := range nums {
		rem := target - num
		if _, ok := seen[rem]; ok {
			out[0] = seen[rem]
			out[1] = i
			break
		}

		seen[num] = i
	}

	return out
}
