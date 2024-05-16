package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(5))
}

func hammingWeight(n int) int {
	var out int

	for n != 0 {
		if n&1 == 1 {
			out++
		}

		n >>= 1
	}

	return out
}
