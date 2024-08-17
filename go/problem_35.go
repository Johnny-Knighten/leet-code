package main

func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		midpoint := (i + j) / 2

		switch {
		case nums[midpoint] == target:
			return midpoint
		case nums[midpoint] > target:
			j = midpoint - 1
		default:
			i = midpoint + 1
		}
	}

	return i
}
