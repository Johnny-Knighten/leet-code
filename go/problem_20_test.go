package main

import "testing"

func TestProblem20(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := "()"
		expected := true

		result := isValid(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := "()[]{}"
		expected := true

		result := isValid(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := "(]"
		expected := false

		result := isValid(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		input := "(("
		expected := false

		result := isValid(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
