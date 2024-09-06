class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        patternCharToWordMap = {}
        s = s.split()
        if len(pattern) != len(s):
            return False

        for i in range(len(pattern)):
            if pattern[i] in patternCharToWordMap:
                if patternCharToWordMap[pattern[i]] != s[i]:
                    return False
            else:
                patternCharToWordMap[pattern[i]] = s[i]

        return len(set(patternCharToWordMap.keys())) == len(set(patternCharToWordMap.values()))


if __name__ == "__main__":
    s = Solution()
    assert s.wordPattern("abba", "dog cat cat dog"), "Test case 1 failed"
    assert not s.wordPattern("abba", "dog cat cat fish"), "Test case 2 failed"
    assert not s.wordPattern("aaaa", "dog cat cat dog"), "Test case 3 failed"
    assert not s.wordPattern("abba", "dog dog dog dog"), "Test case 4 failed"
    assert s.wordPattern("abc", "b c a"), "Test case 5 failed"
    print("All test cases passed")
