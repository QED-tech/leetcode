package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingDistance(4, 10))
}

func hammingDistance(x int, y int) int {
	z := x ^ y

	out := 0

	for z != 0 {
		if z&1 == 1 {
			out++
		}

		z >>= 1
	}

	return out
}
