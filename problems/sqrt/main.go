package main

import "fmt"

func main() {
	r := mySqrt(8)
	fmt.Println(r)
}

func mySqrt(x int) int {
	var multiplier int

	for {
		if (multiplier * multiplier) > x/2 {
			return multiplier - 1
		}
		multiplier++
	}
}
