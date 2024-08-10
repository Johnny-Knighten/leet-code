package main

import "testing"

func TestProblem9(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := 121
		expected := true

		result := isPalindrome(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := -121
		expected := false

		result := isPalindrome(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := 10
		expected := false

		result := isPalindrome(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})
}
