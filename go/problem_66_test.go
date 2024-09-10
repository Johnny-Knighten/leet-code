package main

import (
	"slices"
	"testing"
)

func TestProblem66(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{1, 2, 4}

		result := plusOne(input)

		if !slices.Equal(expected, result) {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{4, 3, 2, 1}
		expected := []int{4, 3, 2, 2}

		result := plusOne(input)

		if !slices.Equal(expected, result) {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := []int{9}
		expected := []int{1, 0}

		result := plusOne(input)

		if !slices.Equal(expected, result) {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})
}
