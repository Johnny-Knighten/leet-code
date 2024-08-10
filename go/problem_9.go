package main

import "strconv"

func isPalindrome(x int) bool {
	xAsStr := strconv.Itoa(x)
	strLen := len(xAsStr)
	for i := 0; i < strLen/2; i++ {
		if xAsStr[i] != xAsStr[strLen-1-i] {
			return false
		}
	}
	return true
}
