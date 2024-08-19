from typing import List


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        result = 0
        for num in nums:
            result = result ^ num
        return result


if __name__ == "__main__":
    s = Solution()
    assert s.singleNumber([2, 2, 1]) == 1, "Test case 1 failed"
    assert s.singleNumber([4, 1, 2, 1, 2]) == 4, "Test case 2 failed"
    assert s.singleNumber([1]) == 1, "Test case 3 failed"
    print("All test cases passed")
