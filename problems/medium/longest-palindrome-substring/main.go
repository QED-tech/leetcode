package main

import (
	"fmt"
)

func main() {
	r := longestPalindrome("cbbd")

	fmt.Println(r)
}

func longestPalindrome(s string) string {
	var max int
	var res string

	for i := 0; i < len(s); i++ {
		if pal := expandFromCenter(s, i, i+1); pal != "" {
			if len(pal) > max {
				max = len(pal)
				res = pal
			}
		}
		if pal := expandFromCenter(s, i, i+1); pal != "" {
			if len(pal) > max {
				max = len(pal)
				res = pal
			}
		}
	}

	return res
}

func expandFromCenter(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]
}
