class Solution:
    def reverseBits(self, n: int) -> int:
        n = (n & 0xffff0000) >> 16 | (n & 0x0000ffff) << 16
        n = (n & 0xff00ff00) >> 8 | (n & 0x00ff00ff) << 8
        n = (n & 0xf0f0f0f0) >> 4 | (n & 0x0f0f0f0f) << 4
        n = (n & 0xCCCCCCCC) >> 2 | (n & 0x33333333) << 2
        n = (n & 0xAAAAAAAA) >> 1 | (n & 0x55555555) << 1
        return n


if __name__ == "__main__":
    s = Solution()
    assert s.reverseBits(43261596) == 964176192, "Test case 1 failed"
    assert s.reverseBits(4294967293) == 3221225471, "Test case 2 failed"
    print("All test cases passed")
