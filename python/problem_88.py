from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        i = m -1
        j = n -1
        k = n + m -1
        while j >= 0:
            if i >= 0 and nums1[i] > nums2[j]:
                nums1[k] = nums1[i]
                i = i - 1
            else:
                nums1[k] = nums2[j]
                j = j - 1
            k = k - 1

if __name__ == "__main__":
    s = Solution()
    destList = [1, 2, 3, 0, 0, 0]
    s.merge(destList, 3, [2, 5, 6], 3)
    assert destList == [1, 2, 2, 3, 5, 6], "Test case 1 failed"

    destList = [1]
    s.merge(destList, 1, [], 0)
    assert destList == [1], "Test case 2 failed"

    destList = [0]
    s.merge(destList, 0, [1], 1)
    assert destList == [1], "Test case 3 failed"

    print("All test cases passed")
