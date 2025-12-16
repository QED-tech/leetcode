package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convertToTitle(2002))
}

func convertToTitle(columnNumber int) string {
	var (
		stack []rune
		b     strings.Builder
	)

	for columnNumber > 0 {
		columnNumber--

		symbol := letterByIndex(columnNumber % 26)

		stack = append(stack, symbol)

		columnNumber /= 26
	}

	for i := len(stack) - 1; i >= 0; i-- {
		b.WriteRune(stack[i])
	}

	return b.String()
}

func letterByIndex(n int) rune {
	if n < 0 || n > 25 {
		return 0
	}
	return rune('A' + n)
}
