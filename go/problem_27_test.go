package main

import (
	"slices"
	"testing"
)

func TestProblem27(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{3, 2, 2, 3}
		inputVal := 3
		expectedK := 2
		expectedSlice := []int{2, 2}

		result := removeElement(input, inputVal)

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
		input := []int{0, 1, 2, 2, 3, 0, 4, 2}
		inputVal := 2
		expectedK := 5
		expectedSlice := []int{0, 1, 4, 0, 3}

		result := removeElement(input, inputVal)

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
