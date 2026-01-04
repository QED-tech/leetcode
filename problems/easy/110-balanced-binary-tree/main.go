package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}

	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}

	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 4}

	result := isBalanced(root)
	fmt.Println(result)
}

func isBalanced(root *TreeNode) bool {
	return balanced(root) != -1
}

func balanced(root *TreeNode) float64 {
	if root == nil {
		return 0
	}

	leftHeight := balanced(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := balanced(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if math.Abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}
