package main

import "testing"

func TestProblem58(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := "Hello World"
		expected := 5

		result := lengthOfLastWord(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		input := "   fly me   to   the moon  "
		expected := 4

		result := lengthOfLastWord(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		input := "luffy is still joyboy"
		expected := 6

		result := lengthOfLastWord(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v, Got: %v", expected, result)
		}
	})
}
