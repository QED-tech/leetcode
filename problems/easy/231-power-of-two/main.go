package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(15))
}

func isPowerOfTwo(num int) bool {
	return num > 0 && (num&(num-1)) == 0
}
