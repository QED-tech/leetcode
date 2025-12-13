package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 5}

	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}

	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}

	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}

	root.Right.Right.Right = &TreeNode{Val: 1}

	fmt.Println(dfs(root, 22))
}

func dfs(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	if dfs(root.Left, targetSum) {
		return true
	}

	if dfs(root.Right, targetSum) {
		return true
	}

	return false
}
