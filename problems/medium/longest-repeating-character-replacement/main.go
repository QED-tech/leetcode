package main

import "fmt"

func main() {
	r := characterReplacement("AABA", 1)
	fmt.Println(r)
}

func characterReplacement(s string, k int) int {
	var (
		left, maxChars, ans int
	)

	store := make(map[string]int)

	for right := 0; right < len(s); right++ {
		store[string(s[right])]++
		maxChars = max(maxChars, store[string(s[right])])

		for (right-left+1)-maxChars > k {
			store[string(s[left])]--
			left++
		}

		ans = max(right - left + 1)
	}

	return ans
}
