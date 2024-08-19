package main

import (
	"testing"
)

func TestProblem136(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{4, 1, 2, 1, 2}
		expected := 4

		result := singleNumber(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{2, 2, 1}
		expected := 1

		result := singleNumber(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{1}
		expected := 1

		result := singleNumber(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
