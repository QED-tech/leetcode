package tools

import (
	"fmt"
	"strings"
)

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

func (q *PriorityQ) Print() {
	var sb strings.Builder

	current := q.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("{v-%d,p-%d}->", current.val, current.priority))
		current = current.next
	}

	fmt.Println(sb.String())
}
