package main

import "testing"

func TestProblem1(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{2, 7, 11, 15}
		target := 9
		expected := []int{0, 1}

		result := twoSum(input, target)

		if result == nil {
			t.Error("Expected a result, not nil")
		}

		if (result[0] != expected[0] && result[1] != expected[1]) && (result[0] != expected[1] && result[1] != expected[0]) {
			t.Errorf("Got incorrect result. Expected: %v, %v. Got: %v, %v", expected[0], expected[1], result[0], result[1])
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{3, 2, 4}
		target := 6
		expected := []int{1, 2}

		result := twoSum(input, target)

		if result == nil {
			t.Error("Expected a result, not nil")
		}

		if (result[0] != expected[0] && result[1] != expected[1]) && (result[0] != expected[1] && result[1] != expected[0]) {
			t.Errorf("Got incorrect result. Expected: %v, %v. Got: %v, %v", expected[0], expected[1], result[0], result[1])
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := []int{3, 3}
		target := 6
		expected := []int{0, 1}

		result := twoSum(input, target)

		if result == nil {
			t.Error("Expected a result, not nil")
		}

		if (result[0] != expected[0] && result[1] != expected[1]) && (result[0] != expected[1] && result[1] != expected[0]) {
			t.Errorf("Got incorrect result. Expected: %v, %v. Got: %v, %v", expected[0], expected[1], result[0], result[1])
		}
	})

}
