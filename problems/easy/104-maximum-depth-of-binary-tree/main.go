package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 33}
	root.Left.Left = &TreeNode{Val: 3}

	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Left.Left = &TreeNode{Val: 5}

	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
}
