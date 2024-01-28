package main

import (
	"fmt"
)

func main() {
	r := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(r)
}

func lengthOfLongestSubstring(s string) int {
	var (
		left,
		maxLen int
	)

	store := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		store[s[right]]++

		if charCount := store[s[right]]; charCount > 1 {
			for {
				if charCount := store[s[right]]; charCount <= 1 {
					break
				}
				store[s[left]]--
				left++
			}
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
