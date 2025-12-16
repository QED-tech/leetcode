package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3749))
}

/*
MMMDCCXLIX

	1. 3449
	   MMM

	2. 449
	   MMMCD

	3. 49
*/

func intToRoman(num int) string {
	var (
		b strings.Builder
	)

loop:
	for num > 0 {
		switch {
		case num >= 1000:
			b.WriteString("M")
			num -= 1000
		case num >= 900:
			b.WriteString("CM")
			num -= 900
		case num >= 500:
			b.WriteString("D")
			num -= 500
		case num >= 400:
			b.WriteString("CD")
			num -= 400
		case num >= 100:
			b.WriteString("C")
			num -= 100
		case num >= 90:
			b.WriteString("XC")
			num -= 90
		case num >= 50:
			b.WriteString("L")
			num -= 50
		case num >= 40:
			b.WriteString("XL")
			num -= 40
		case num >= 10:
			b.WriteString("X")
			num -= 10
		case num >= 9:
			b.WriteString("IX")
			num -= 9
		case num >= 5:
			b.WriteString("V")
			num -= 5
		case num >= 4:
			b.WriteString("IV")
			num -= 4
		case num >= 1:
			b.WriteString("I")
			num -= 1
		default:
			break loop
		}
	}

	return b.String()
}
