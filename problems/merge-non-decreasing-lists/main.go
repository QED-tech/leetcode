package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{-2, 3, 3}
	b := []int{-5}

	r := mergeSorted(a, b)
	fmt.Println(r)
}

// [-2,3,3] [-5,0]
// [-5, -2, 0, 3, 3]
var max = math.MaxInt

func mergeSorted(a, b []int) []int {
	out := make([]int, 0)

	a = append(a, max)
	b = append(b, max)

	p1, p2 := 0, 0

	for a[p1] != max || b[p2] != max {
		if a[p1] <= b[p2] {
			out = append(out, a[p1])
			p1++
		} else {
			out = append(out, b[p2])
			p2++
		}
	}

	return out
}
