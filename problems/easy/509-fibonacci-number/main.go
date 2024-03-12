package main

import "fmt"

func main() {
	fmt.Println(fib(3))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	store := [2]int{0, 1}

	for i := 2; i < n; i++ {
		sum := store[0] + store[1]
		store[0] = store[1]
		store[1] = sum
	}

	return store[0] + store[1]
}
