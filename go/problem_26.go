package main

func removeDuplicates(nums []int) int {
	nextInsertPos := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nextInsertPos] = nums[i]
			nextInsertPos++
		}
	}
	return nextInsertPos
}
