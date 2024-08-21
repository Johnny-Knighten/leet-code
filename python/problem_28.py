class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if len(needle) > len(haystack):
            return -1
        
        for i in range(0, len(haystack)-len(needle)+1):
            if needle == haystack[i:i+len(needle)]:
                return i
        
        return -1
    
if __name__ == "__main__":
    s = Solution()
    assert s.strStr(haystack = "sadbutsad", needle = "sad") == 0, "Test case 1 failed"
    assert s.strStr(haystack = "leetcode", needle = "leeto") == -1, "Test case 2 failed"
    print("All test cases passed")
