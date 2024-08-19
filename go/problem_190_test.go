package main

import (
	"testing"
)

func TestProblem190(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		var input uint32 = 43261596
		var expected uint32 = 964176192

		result := reverseBits(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		var input uint32 = 4294967293
		var expected uint32 = 3221225471

		result := reverseBits(input)

		if result != expected {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, result)
		}
	})

}
