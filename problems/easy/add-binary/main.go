package main

import (
	"fmt"
	"leetcode/tools"
	"strconv"
	"strings"
)

func main() {
	r := addBinary("1", "111")

	fmt.Println(r)
}

func addBinary(a string, b string) string {
	normalizeA, normalizeB := a, b

	if len(a) > len(b) {
		normalizeB = normalizer(len(a)-len(b), b)
	} else {
		normalizeA = normalizer(len(b)-len(a), a)
	}

	var res strings.Builder
	var carry, remainder int

	for i := len(normalizeA) - 1; i >= 0; i-- {
		d1, _ := strconv.Atoi(string(normalizeA[i]))
		d2, _ := strconv.Atoi(string(normalizeB[i]))

		numerator := (d1 + d2 + carry)

		carry, remainder = numerator/2, numerator%2
		res.WriteString(strconv.Itoa(remainder))
	}

	if carry > 0 {
		res.WriteString(strconv.Itoa(carry))
	}

	return tools.ReverseString(res.String())
}

func normalizer(len int, str string) string {
	var r strings.Builder

	i := len
	for i > 0 {
		r.WriteRune('0')
		i--
	}

	r.WriteString(str)

	return r.String()
}
