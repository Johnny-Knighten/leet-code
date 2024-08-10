class Solution:
    def isPalindrome(self, x: int) -> bool:
        forward = str(x)
        backward = forward[::-1]
        return forward == backward
    
if __name__ == "__main__":
    s = Solution()
    assert s.isPalindrome(121) == True, "Test case 1 failed"
    assert s.isPalindrome(-121) == False, "Test case 1 failed"
    assert s.isPalindrome(10) == False, "Test case 1 failed"
    print("All test cases passed")
