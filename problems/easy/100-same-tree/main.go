package main

import (
	"fmt"
	"leetcode/tools"
)

/*
	1
   / \
  2	 3

[1,2,3]


*/

func main() {
	p := &tools.TreeNode{Val: 1}
	p.Left = &tools.TreeNode{Val: 2}
	p.Right = &tools.TreeNode{Val: 3}

	q := &tools.TreeNode{Val: 1}
	q.Left = &tools.TreeNode{Val: 2}
	q.Right = &tools.TreeNode{Val: 3}

	fmt.Println(isSameTree(p, q))
}
func isSameTree(p *tools.TreeNode, q *tools.TreeNode) bool {
	if p == nil && q != nil {
		return false
	}

	if q == nil && p != nil {
		return false
	}

	if q == nil && p == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
