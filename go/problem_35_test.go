package main

import "testing"

func TestProblem35(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{1, 3, 5, 6}
		target := 5
		expected := 2

		result := searchInsert(input, target)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{1, 3, 5, 6}
		target := 2
		expected := 1

		result := searchInsert(input, target)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := []int{1, 3, 5, 6}
		target := 7
		expected := 4

		result := searchInsert(input, target)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		input := []int{1, 3}
		target := 0
		expected := 0

		result := searchInsert(input, target)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
