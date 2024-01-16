package main

import (
	"fmt"
)

//5
// comb = ceil(5 / 2)

// 1 + 1 + 1 + 1 + 1

// 2 + 1 + 1 + 1
// 1 + 2 + 1 + 1
// 1 + 1 + 2 + 1
// 1 + 1 + 1 + 2

// 1 + 2 + 2
// 2 + 1  + 2
// 2 + 2 + 1

func main() {
	r := climbStairs(5)
	fmt.Println(r)
}

func climbStairs(n int) int {
	if n < 1 {
		return n
	}

	var res int
	fib := []int{0, 1}

	for i := 1; i <= n; i++ {
		res = fib[0] + fib[1]
		fib[0] = fib[1]
		fib[1] = res
	}

	return res
}
