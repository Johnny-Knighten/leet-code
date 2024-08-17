from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        i = 0
        j = len(nums)-1
        while i <= j:
            midpoint = (i+j) // 2

            if nums[midpoint] == target:
                return midpoint
            elif nums[midpoint] < target:
                i = midpoint + 1
            else:
                j = midpoint - 1

        return i


if __name__ == "__main__":
    s = Solution()
    assert s.searchInsert([1, 3, 5, 6], 5) == 2, "Test case 1 failed"
    assert s.searchInsert([1, 3, 5, 6], 2) == 1, "Test case 2 failed"
    assert s.searchInsert([1, 3, 5, 6], 7) == 4, "Test case 3 failed"
    assert s.searchInsert([1, 3], 0) == 0, "Test case 4 failed"
    print("All test cases passed")
