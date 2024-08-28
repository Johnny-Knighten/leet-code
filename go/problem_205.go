package main

func isIsomorphic(s string, t string) bool {
	sCharLastIndexMap := map[rune]int{}
	tCharLastIndexMap := map[rune]int{}

	for i := range len(s) {
		sCharLastIndex, tCharLastIndex := 0, 0
		if sVal, ok := sCharLastIndexMap[rune(s[i])]; ok {
			sCharLastIndex = sVal
		}

		if tVal, ok := tCharLastIndexMap[rune(t[i])]; ok {
			tCharLastIndex = tVal
		}

		if sCharLastIndex != tCharLastIndex {
			return false
		}

		sCharLastIndexMap[rune(s[i])] = i + 1
		tCharLastIndexMap[rune(t[i])] = i + 1

	}

	return true
}
