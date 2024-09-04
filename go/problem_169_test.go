package main

import "testing"

func TestProblem169(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{3, 2, 3}
		expected := 3

		result := majorityElement(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{2, 2, 1, 1, 1, 2, 2}
		expected := 2

		result := majorityElement(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
