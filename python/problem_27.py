from typing import List

class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        start = 0
        end = len(nums)-1
        count = 0

        while start <= end:
            if nums[start] == val:
              nums[start], nums[end] = nums[end], nums[start]
              end = end -1
            else:
                start = start + 1


        return start

if __name__ == "__main__":
    s = Solution()
    input = [3,2,2,3]
    result = s.removeElement(input, 3)
    assert result == 2 and sorted(input[:result]) == sorted([2,2]), "Test case 1 failed"
    input = [0,1,2,2,3,0,4,2]
    result = s.removeElement(input, 2)
    assert result == 5 and sorted(input[:result]) == sorted([0,1,4,0,3]), "Test case 2 failed"
    print("All test cases passed")