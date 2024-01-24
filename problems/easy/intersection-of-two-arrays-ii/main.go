package main

import "fmt"

func main() {
	r := intersect([]int{4, 9, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(r)
}

func intersect(nums1 []int, nums2 []int) []int {
	seen := make(map[int]int)
	var out []int

	for _, n := range nums1 {
		seen[n]++
	}

	for _, n := range nums2 {
		if _, ok := seen[n]; ok {
			if seen[n] > 0 {
				out = append(out, n)
			}
		}
		seen[n]--
	}

	return out
}
