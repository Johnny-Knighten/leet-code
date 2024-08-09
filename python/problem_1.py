from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numbersSeen = dict()
        for i in range(len(nums)):
            requiredNum = target - nums[i]
            if requiredNum in numbersSeen:
                return [i, numbersSeen[requiredNum]]
            numbersSeen[nums[i]] = i
        return []


if __name__ == "__main__":
    s = Solution()
    assert set(s.twoSum([2, 7, 11, 15], 9)) == set(
        [0, 1]), "Test case 1 failed"
    assert set(s.twoSum([3, 2, 4], 6)) == set([1, 2]), "Test case 2 failed"
    assert set(s.twoSum([3, 3], 6)) == set([0, 1]), "Test case 3 failed"
    print("All test cases passed")
