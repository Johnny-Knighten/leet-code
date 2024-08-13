package main

import "testing"

func TestProblem14(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []string{"flower", "flow", "flight"}
		expected := "fl"

		result := longestCommonPrefix(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []string{"dog", "racecar", "car"}
		expected := ""

		result := longestCommonPrefix(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := []string{"cir", "car"}
		expected := "c"

		result := longestCommonPrefix(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		input := []string{"ab", "a"}
		expected := "a"

		result := longestCommonPrefix(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 5", func(t *testing.T) {
		input := []string{"abab", "aba", ""}
		expected := ""

		result := longestCommonPrefix(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
