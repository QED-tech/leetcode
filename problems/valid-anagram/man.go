package main

import "fmt"

func main() {
	r := isAnagram("anagram", "nagaram")
	fmt.Println(r)
}

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	for _, char := range t {
		m[char]--
	}

	for k := range m {
		if m[k] != 0 {
			return false
		}
	}

	return true
}
