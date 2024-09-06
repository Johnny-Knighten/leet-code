package main

import "testing"

func TestProblem290(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input1 := "abba"
		input2 := "dog cat cat dog"
		expected := true

		result := wordPattern(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input1 := "abba"
		input2 := "dog cat cat fish"
		expected := false

		result := wordPattern(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input1 := "aaaa"
		input2 := "dog cat cat dog"
		expected := false

		result := wordPattern(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		input1 := "abba"
		input2 := "dog dog dog dog"
		expected := false

		result := wordPattern(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 5", func(t *testing.T) {
		input1 := "abc"
		input2 := "b c a"
		expected := true

		result := wordPattern(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
