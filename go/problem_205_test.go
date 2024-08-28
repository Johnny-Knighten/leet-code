package main

import "testing"

func TestProblem205(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input1 := "egg"
		input2 := "add"
		expected := true

		result := isIsomorphic(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input1 := "foo"
		input2 := "bar"
		expected := false

		result := isIsomorphic(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input1 := "paper"
		input2 := "title"
		expected := true

		result := isIsomorphic(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		input1 := "bbbaaaba"
		input2 := "aaabbbba"
		expected := false

		result := isIsomorphic(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
