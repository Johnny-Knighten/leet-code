package main

import "testing"

func TestProblem28(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input1 := "sadbutsad"
		input2 := "sad"
		expected := 0

		result := strStr(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input1 := "leetcode"
		input2 := "leeto"
		expected := -1

		result := strStr(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
