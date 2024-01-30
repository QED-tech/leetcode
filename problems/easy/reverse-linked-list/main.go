package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1 -> 2 -> 3 -> 4 -> 5 -> 10

1 <- 2 <- 3 <- 4 <- 5 <- 10
*/

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 10,
			},
		},
	}

	fmt.Println(reverseList(node))
}

func reverseList(head *ListNode) *ListNode {
	var reversed *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = reversed
		reversed = head
		head = tmp
	}

	return reversed
}
