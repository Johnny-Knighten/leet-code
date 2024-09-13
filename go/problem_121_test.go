package main

import (
	"testing"
)

func TestProblem121(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{7, 1, 5, 3, 6, 4}
		expected := 5

		result := maxProfit(input)

		if expected != result {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{7, 6, 4, 3, 1}
		expected := 0

		result := maxProfit(input)

		if expected != result {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
