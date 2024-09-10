from typing import List


class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        digits[-1] += 1
        for i in range(len(digits)-1, -1, -1):
            if digits[i] == 10:
                digits[i] = 0
                if i == 0:
                    digits.insert(0, 1)
                else:
                    digits[i-1] += 1
        return digits


if __name__ == "__main__":
    s = Solution()
    assert s.plusOne([1, 2, 3]) == [1, 2, 4], "Test case 1 failed"
    assert s.plusOne([4, 3, 2, 1]) == [4, 3, 2, 2], "Test case 2 failed"
    assert s.plusOne([9]) == [1, 0], "Test case 3 failed"
    print("All test cases passed")
