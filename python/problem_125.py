import re


class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = re.sub(r'[\W_+]', '', s).lower().strip()
        start = 0
        end = len(s)-1

        while start < end:
            if s[start] == s[end]:
                start += 1
                end -= 1
            else:
                return False

        return True


if __name__ == "__main__":
    s = Solution()
    assert s.isPalindrome(
        "A man, a plan, a canal: Panama") == True, "Test case 1 failed"
    assert s.isPalindrome("race a car") == False, "Test case 2 failed"
    assert s.isPalindrome(" ") == True, "Test case 3 failed"
    assert s.isPalindrome("ab_a") == True, "Test case 4 failed"
    print("All test cases passed")
