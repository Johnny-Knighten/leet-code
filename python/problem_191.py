class Solution:
    def hammingWeight(self, n: int) -> int:
        count = 0
        while n != 0:
            if (n & 1):
                count += 1
            n >>= 1
        return count


if __name__ == "__main__":
    s = Solution()
    assert s.hammingWeight(11) == 3, "Test case 1 failed"
    assert s.hammingWeight(128) == 1, "Test case 2 failed"
    assert s.hammingWeight(2147483645) == 30, "Test case 3 failed"
    print("All test cases passed")
