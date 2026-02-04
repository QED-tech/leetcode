package main

import "fmt"

func getConcatenation(nums []int) []int {
	l := len(nums)

	ans := make([]int, l*2)

	copy(ans[0:l], nums)
	copy(ans[l:], nums)

	return ans
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
}
