package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

type stackItem struct {
	path  []int
	sum   int
	start int
}

type stack []stackItem

func (s *stack) pop() stackItem {
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}

func (s *stack) push(el stackItem) {
	*s = append(*s, el)
}

func combinationSum(candidates []int, target int) [][]int {
	var out [][]int

	var stack stack
	stack.push(stackItem{})

	for len(stack) > 0 {
		el := stack.pop()

		if el.sum > target {
			continue
		}

		if el.sum == target {
			out = append(out, el.path)
			continue
		}

		for i := el.start; i < len(candidates); i++ {
			nextPath := make([]int, len(el.path)+1)
			copy(nextPath, el.path)
			nextPath[len(el.path)] = candidates[i]

			stack.push(stackItem{
				start: i,
				sum:   el.sum + candidates[i],
				path:  nextPath,
			})
		}

	}

	return out
}
