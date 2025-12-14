package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right.Right = &TreeNode{Val: 8}

	root.Left.Right.Left = &TreeNode{Val: 6}
	root.Left.Right.Right = &TreeNode{Val: 7}

	root.Right.Right.Left = &TreeNode{Val: 9}

	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var (
		out []int
		dfs func(node *TreeNode)
	)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		out = append(out, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return out
}
