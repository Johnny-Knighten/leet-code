package main

import "testing"

func TestProblem125(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := "A man, a plan, a canal: Panama"
		expected := true

		result := isPalindrome125(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := "race a car"
		expected := false

		result := isPalindrome125(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := " "
		expected := true

		result := isPalindrome125(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		input := "ab_a"
		expected := true

		result := isPalindrome125(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
