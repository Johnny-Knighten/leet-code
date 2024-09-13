from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0
        currentLow = prices[0]
        for i in range(1, len(prices)):
            if prices[i] < currentLow:
                currentLow = prices[i]
            elif prices[i]-currentLow > profit:
                profit = prices[i]-currentLow
        return profit


if __name__ == "__main__":
    s = Solution()
    assert s.maxProfit([7, 1, 5, 3, 6, 4]) == 5, "Test case 1 failed"
    assert s.maxProfit([7, 6, 4, 3, 1]) == 0, "Test case 2 failed"
    print("All test cases passed")
