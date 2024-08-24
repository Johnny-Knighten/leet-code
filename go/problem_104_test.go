package main

import "testing"

func TestProblem104(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
		expected := 3

		result := maxDepth(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
		expected := 2

		result := maxDepth(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
