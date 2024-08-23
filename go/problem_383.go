package main

func canConstruct(ransomNote string, magazine string) bool {
	ransomMap := map[rune]int{}
	for _, char := range ransomNote {
		if val, ok := ransomMap[char]; ok {
			ransomMap[char] = val + 1
		} else {
			ransomMap[char] = 1
		}
	}

	magazineMap := map[rune]int{}
	for _, char := range magazine {
		if val, ok := magazineMap[char]; ok {
			magazineMap[char] = val + 1
		} else {
			magazineMap[char] = 1
		}
	}

	for key, val := range ransomMap {
		if magazineVal, ok := magazineMap[key]; !ok || val > magazineVal {
			return false
		}
	}

	return true
}
