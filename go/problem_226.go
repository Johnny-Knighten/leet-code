package main

// See Problem 104 For Tree Node Implementation
func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	root.Left, root.Right = root.Right, root.Left

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)

	return root
}
