package main

import (
	"strings"
	"unicode"
)

func isPalindrome125(s string) bool {

	cleanStrRunes := []rune{}
	for _, letter := range s {
		if unicode.IsNumber(letter) || unicode.IsLetter(letter) {
			cleanStrRunes = append(cleanStrRunes, letter)
		}
	}

	cleanStr := strings.ToLower(string(cleanStrRunes))

	start := 0
	end := len(cleanStr) - 1
	for start < end {
		if cleanStr[start] != cleanStr[end] {
			return false
		} else {
			start++
			end--
		}
	}

	return true
}
