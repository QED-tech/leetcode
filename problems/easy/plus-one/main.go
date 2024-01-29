package main

import "fmt"

/*

You are given a large integer represented as an integer array digits,
where each digits[i] is the ith digit of the integer.
The digits are ordered from most significant to least significant in left-to-right order.
The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.


*/

func main() {
	r := plusOne([]int{9, 9, 9})
	fmt.Println(r) //1000
}

func plusOne(digits []int) []int {
	var out []int
	remainder := 1

	for i := len(digits) - 1; i >= 0; i-- {
		sum := remainder + digits[i]
		out = append(out, sum%10)
		remainder = sum / 10
	}

	if remainder == 1 {
		out = append(out, 1)
	}

	reverse(out)
	return out
}

func reverse(digits []int) {
	for i, j := 0, len(digits)-1; j > i; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
}
