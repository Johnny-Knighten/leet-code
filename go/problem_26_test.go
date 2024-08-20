package main

import (
	"slices"
	"testing"
)

func TestProblem26(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{1, 1, 2}
		expectedK := 2
		expectedSlice := []int{1, 2}

		result := removeDuplicates(input)

		if result != expectedK {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expectedK, result)
		}

		slices.Sort(expectedSlice)
		slices.Sort(input[:result])
		if !slices.Equal(input[:result], expectedSlice) {
			t.Errorf("Got incorrect slice. Expected: %v. Got: %v", input[:result], expectedSlice)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		expectedK := 5
		expectedSlice := []int{0, 1, 2, 3, 4}

		result := removeDuplicates(input)

		if result != expectedK {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expectedK, result)
		}

		slices.Sort(expectedSlice)
		slices.Sort(input[:result])
		if !slices.Equal(input[:result], expectedSlice) {
			t.Errorf("Got incorrect slice. Expected: %v. Got: %v", input[:result], expectedSlice)
		}
	})

}
