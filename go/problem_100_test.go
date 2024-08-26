package main

import "testing"

func TestProblem100(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
		input2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
		expected := true

		result := isSameTree(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
		input2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
		expected := false

		result := isSameTree(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
		input2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
		expected := false

		result := isSameTree(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
