package main

import (
	"slices"
	"testing"
)

func TestProblem88(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums1Size := 3
		nums2 := []int{2, 5, 6}
		nums2Size := 3

		expected := []int{1, 2, 2, 3, 5, 6}

		merge(nums1, nums1Size, nums2, nums2Size)

		if !slices.Equal(nums1, expected) {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, nums1)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		nums1 := []int{1}
		nums1Size := 1
		nums2 := []int{}
		nums2Size := 0

		expected := []int{1}

		merge(nums1, nums1Size, nums2, nums2Size)

		if !slices.Equal(nums1, expected) {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, nums1)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		nums1 := []int{0}
		nums1Size := 0
		nums2 := []int{1}
		nums2Size := 1

		expected := []int{1}

		merge(nums1, nums1Size, nums2, nums2Size)

		if !slices.Equal(nums1, expected) {
			t.Errorf("Got incorrect result. Expected: %v. Got: %v", expected, nums1)
		}
	})

}
