package tools

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintLinkedList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
