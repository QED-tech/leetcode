package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextGreaterElements([]int{3, 8, 4, 1, 2}))
}

type stack []int

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() int {
	el := s.top()
	*s = (*s)[:len(*s)-1]
	return el
}

func (s *stack) push(el int) {
	*s = append(*s, el)
}

func nextGreaterElements(nums []int) []int {
	out := make([]int, len(nums))
	for i := range out {
		out[i] = -1
	}

	s := make([]int, 0, len(nums))
	stack := stack(s)
	for i := 0; i < 2*len(nums); i++ {
		i := i % len(nums)

		for len(stack) != 0 && nums[i] > nums[stack.top()] {
			idx := stack.pop()
			out[idx] = nums[i]
		}

		stack.push(i)

		fmt.Println(stack)
	}

	return out
}
