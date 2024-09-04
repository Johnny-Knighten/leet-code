from typing import List


class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        currentLeader = 0
        count = 0
        for num in nums:
            if count == 0:
                currentLeader = num
                count = 1
            elif num == currentLeader:
                count += 1
            else:
                count -= 1
        return currentLeader


if __name__ == "__main__":
    s = Solution()
    assert s.majorityElement([3, 2, 3]) == 3, "Test case 1 failed"
    assert s.majorityElement([2, 2, 1, 1, 1, 2, 2]) == 2, "Test case 2 failed"
    print("All test cases passed")
