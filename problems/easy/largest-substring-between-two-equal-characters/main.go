package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("aa"))
}

/*
l r
b w  e  a  b  b  c  a
*/

func maxLengthBetweenEqualCharacters(s string) int {
	maxLen := -1
	store := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if index, ok := store[s[i]]; ok {
			maxLen = max(maxLen, i-index-1)
		}

		store[s[i]] = i
	}

	return maxLen
}
