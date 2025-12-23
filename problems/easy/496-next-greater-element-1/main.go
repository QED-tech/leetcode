package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var (
		out   []int
		left  int
		n     = len(nums1)
		right = n
	)

	nums1 = append(nums1, nums2...)

	for left < right {
		var (
			sawSearchElem bool
			foundGreater  bool
		)

		for right < len(nums1) {

			if nums1[right] == nums1[left] {
				sawSearchElem = true
			}

			if sawSearchElem && nums1[right] > nums1[left] {
				out = append(out, nums1[right])
				foundGreater = true
				break
			}
			right++
		}

		if !foundGreater {
			out = append(out, -1)
		}

		right = n
		left++
	}

	return out
}
