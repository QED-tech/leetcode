package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

/*
  l             r
1 8 6 2 5 4 8 3 7
  _______________

max = 0
l = 0
r = len(height) - 1

while l < r {
	area := min(l, r) * (r - l)

	max = max(max, area)

	if r > l {
	  l++
	} else {
	  r--
	}
}
*/

func maxArea(height []int) int {
	res := 0
	l := 0
	r := len(height) - 1

	for l < r {
		area := min(height[l], height[r]) * (r - l)

		res = max(res, area)

		if height[r] > height[l] {
			l++
		} else {
			r--
		}

	}

	return res
}
