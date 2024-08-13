package main

func longestCommonPrefix(strs []string) string {
	lcp := strs[0]
	for i := 1; i < len(strs); i++ {
		maxLCPLength := min(len(lcp), len(strs[i]))
		if maxLCPLength == 0 {
			return ""
		}

		for j := 0; j < maxLCPLength; j++ {
			if lcp[j] != strs[i][j] {
				lcp = lcp[0:j]
				break
			} else if j == maxLCPLength-1 {
				lcp = lcp[0 : j+1]
			}
		}
	}
	return lcp
}
