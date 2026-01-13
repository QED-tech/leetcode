package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("([])"))   //true
	fmt.Println(isValid("([)]"))   //false
}

var (
	pairs = map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}
)

func isValid(s string) bool {
	var (
		stack []rune
	)

	for _, r := range s {
		_, isOpen := pairs[r]
		if isOpen {
			stack = append(stack, r)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		open := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pairs[open] != r {
			return false
		}
	}

	return len(stack) == 0
}
