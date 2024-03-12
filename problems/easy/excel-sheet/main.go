package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("ZYA"))
}

/*
	(0*26+26) = 26
	26 * 26 + 25 = 701
	701 * 26 + 1 = 18227
*/

func titleToNumber(columnTitle string) int {
	result := 0
	for _, ch := range columnTitle {
		value := int(ch) - 'A' + 1
		result = result*26 + value
	}

	return result
}
