package main

import "fmt"

func main() {
	r := topKFrequent([]int{1, 3, 2, 1, 2, 1}, 2)
	fmt.Println(r)
}

type PriorityQ struct {
	head *Node
}

func NewPriorityQ() *PriorityQ {
	return &PriorityQ{}
}

type Node struct {
	val      int
	priority int
	next     *Node
}

func (q *PriorityQ) Add(node *Node) {
	if q.head == nil {
		q.head = node
		return
	}

	if node.priority >= q.head.priority {
		q.head = &Node{val: node.val, priority: node.priority, next: q.head}
		return
	}

	current := q.head
	for {
		if current.next == nil {
			current.next = node
			return
		}

		if current.next.priority <= node.priority {
			current.next = &Node{val: node.val, priority: node.priority, next: current.next}
			return
		}

		current = current.next
	}
}

func (q *PriorityQ) Pop() *Node {
	elem := q.head
	q.head = q.head.next

	return elem
}

func topKFrequent(arr []int, k int) []int {
	stor := make(map[int]int)

	for _, num := range arr {
		stor[num]++
	}

	q := NewPriorityQ()
	for key := range stor {
		q.Add(&Node{val: key, priority: stor[key]})
	}

	out := make([]int, k)

	for i := 0; i < k; i++ {
		out[i] = q.Pop().val
	}

	return out
}
