class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        return len(s.strip().split(" ")[-1])


if __name__ == "__main__":
    s = Solution()
    assert s.lengthOfLastWord("Hello World") == 5, "Test case 1 failed"
    assert s.lengthOfLastWord(
        "   fly me   to   the moon  ") == 4, "Test case 2 failed"
    assert s.lengthOfLastWord(
        "luffy is still joyboy") == 6, "Test case 3 failed"
    print("All test cases passed")
