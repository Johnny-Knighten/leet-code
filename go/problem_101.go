package main

// See Problem 104 For Tree Node Implementation
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recurIsSymmetric(root.Left, root.Right)
}

func recurIsSymmetric(leftNode *TreeNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val {
		return false
	}

	return recurIsSymmetric(leftNode.Left, rightNode.Right) && recurIsSymmetric(leftNode.Right, rightNode.Left)
}
