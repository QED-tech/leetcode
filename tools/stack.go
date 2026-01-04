package tools

type stack []int

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() int {
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}

func (s *stack) push(el int) {
	*s = append(*s, el)
}
