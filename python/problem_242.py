class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        sMap = dict()
        for char in s:
            sMap[char] = sMap.get(char, 0) + 1

        for char in t:
            if char in sMap and sMap[char] > 0:
                sMap[char] = sMap[char] - 1
            else:
                return False
        return True


if __name__ == "__main__":
    s = Solution()
    assert s.isAnagram("anagram", "nagaram"), "Test case 1 failed"
    assert not s.isAnagram("rat", "car"), "Test case 2 failed"
    print("All test cases passed")
