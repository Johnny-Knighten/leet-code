package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 1)
}

func dfs(node *TreeNode, currentDepth int) int {
	if node.Left == nil && node.Right == nil {
		return currentDepth
	}

	leftDepth := currentDepth
	rightDepth := currentDepth

	if node.Left != nil {
		leftDepth = dfs(node.Left, currentDepth+1)
	}

	if node.Right != nil {
		rightDepth = dfs(node.Right, currentDepth+1)
	}

	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}
