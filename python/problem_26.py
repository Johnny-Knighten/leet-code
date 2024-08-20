from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        nextInsertPos = 1
        for i in range(1, len(nums)):
            if nums[i] != nums[i-1]:
                nums[nextInsertPos] = nums[i]
                nextInsertPos += 1

        return nextInsertPos


if __name__ == "__main__":
    s = Solution()
    # [1,2,_]
    input = [1, 1, 2]
    result = s.removeDuplicates(input)
    assert result == 2 and input[:result] == [1, 2], "Test case 1 failed"
    # [0,1,2,3,4,_,_,_,_,_]
    input = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
    result = s.removeDuplicates(input)
    assert result == 5 and input[:result] == [
        0, 1, 2, 3, 4], "Test case 2 failed"
    print("All test cases passed")
