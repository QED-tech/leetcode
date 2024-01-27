package main

import "fmt"

func main() {
	r := isValid("()[]{}")
	fmt.Println(r)
}

func isValid(s string) bool {
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		current := string(s[i])
		isClosestParenthes := current == ")" || current == "}" || current == "]"
		if i >= 0 && isClosestParenthes {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			brackets := pop + current

			fmt.Println(brackets)
			if brackets != "()" && brackets != "{}" && brackets != "[]" {
				return false
			}
		} else {
			stack = append(stack, current)
		}
	}

	return true
}
