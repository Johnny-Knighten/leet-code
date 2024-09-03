package main

import "testing"

func TestProblem392(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input1 := "abc"
		input2 := "ahbgdc"
		expected := true

		result := isSubsequence(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input1 := "axc"
		input2 := "ahbgdc"
		expected := false

		result := isSubsequence(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
