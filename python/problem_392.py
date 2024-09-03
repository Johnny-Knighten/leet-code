class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        i, j = 0, 0
        while (i < len(s) and j < len(t)):
            if s[i] == t[j]:
                i += 1
                j += 1
            else:
                j += 1

        return i == len(s)


if __name__ == "__main__":
    s = Solution()
    assert s.isSubsequence("abc", "ahbgdc"), "Test case 1 failed"
    assert not s.isSubsequence("axc", "ahbgdc"), "Test case 2 failed"
    print("All test cases passed")
