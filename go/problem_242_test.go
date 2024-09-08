package main

import "testing"

func TestProblem242(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input1 := "anagram"
		input2 := "nagaram"
		expected := true

		result := isAnagram(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input1 := "rat"
		input2 := "car"
		expected := false

		result := isAnagram(input1, input2)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
