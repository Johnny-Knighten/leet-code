package main

import (
	"testing"
)

func TestProblem191(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := 11
		expected := 3

		result := hammingWeight(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := 128
		expected := 1

		result := hammingWeight(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := 2147483645
		expected := 30

		result := hammingWeight(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
