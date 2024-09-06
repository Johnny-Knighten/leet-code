package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	runes := []rune(pattern)

	if len(words) != len(runes) {
		return false
	}

	patCharToWordMap := make(map[rune]string)
	for i := range len(words) {
		if word, ok := patCharToWordMap[runes[i]]; !ok {
			patCharToWordMap[runes[i]] = words[i]
		} else {
			if word != words[i] {
				return false
			}
		}
	}

	keys := make(map[rune]bool)
	vals := make(map[string]bool)
	for k, v := range patCharToWordMap {
		if _, ok := keys[k]; !ok {
			keys[k] = true
		}

		if _, ok := vals[v]; !ok {
			vals[v] = true
		}
	}

	return len(keys) == len(vals)
}
