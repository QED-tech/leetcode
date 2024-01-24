package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	r := isPalindrome("ab2a")
	fmt.Println(r)
}

func isPalindrome(s string) bool {
	sLen := utf8.RuneCountInString(s)
	start, end := 0, sLen-1

	for end > start {
		if !unicode.IsLetter(rune(s[end])) && !unicode.IsDigit(rune(s[end])) {
			end--
			continue
		}

		if !unicode.IsLetter(rune(s[start])) && !unicode.IsDigit(rune(s[start])) {
			start++
			continue
		}

		if !strings.EqualFold(string(s[start]), string(s[end])) {
			return false
		}

		end--
		start++
	}

	return true
}
