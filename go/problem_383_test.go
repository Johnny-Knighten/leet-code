package main

import "testing"

func TestProblem383(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input1 := "a"
		input2 := "b"
		expected := false

		result := canConstruct(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input1 := "aa"
		input2 := "ab"
		expected := false

		result := canConstruct(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input1 := "aa"
		input2 := "aab"
		expected := true

		result := canConstruct(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
