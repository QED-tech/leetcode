package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,

				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
	}

	deleteDuplicates(h)
	printLinkedList(h)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	h := head

	prev := head
	current := head.Next
	for current != nil {
		next := current.Next

		if prev.Val == current.Val {
			prev.Next = current.Next
		} else {
			prev = current
		}

		current = next
	}

	return h
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
