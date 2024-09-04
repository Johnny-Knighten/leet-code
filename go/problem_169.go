package main

func majorityElement(nums []int) int {
	currentLeder := 0
	leaderCount := 0
	for _, num := range nums {
		if leaderCount == 0 {
			currentLeder = num
			leaderCount = 1
		} else if currentLeder == num {
			leaderCount++
		} else {
			leaderCount--
		}
	}

	return currentLeder
}
