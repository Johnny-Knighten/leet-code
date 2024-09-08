package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)
	for _, charRune := range s {
		if val, ok := sMap[charRune]; ok {
			sMap[charRune] = val + 1
		} else {
			sMap[charRune] = 1
		}
	}

	for _, charRune := range t {
		if val, ok := sMap[charRune]; ok && val > 0 {
			sMap[charRune] = val - 1
		} else {
			return false
		}
	}

	return true
}
