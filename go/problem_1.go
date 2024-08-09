package main

func twoSum(nums []int, target int) []int {
	numbersSeen := map[int]int{}
	for i, num := range nums {
		if val, ok := numbersSeen[target-num]; ok {
			return []int{i, val}
		}
		numbersSeen[num] = i
	}

	return nil
}
