from typing import List

class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        lcp = strs[0]
        for i in range(1, len(strs)):
          maxPossibleLCP = min(len(lcp), len(strs[i]))
          if(maxPossibleLCP == 0):
              return ""
          for j in range(0, maxPossibleLCP):
              if lcp[j] != strs[i][j]:
                  lcp = lcp[0:j]
                  break
              elif j == maxPossibleLCP-1:
                  lcp = lcp[0:j+1]
        return lcp 
            

if __name__ == "__main__":
    s = Solution()
    assert s.longestCommonPrefix(["flower","flow","flight"]) == "fl", "Test case 1 failed"
    assert s.longestCommonPrefix(["dog","racecar","car"]) == "", "Test case 2 failed"
    assert s.longestCommonPrefix(["cir","car"]) == "c", "Test case 3 failed"
    assert s.longestCommonPrefix(["ab","a"]) == "a", "Test case 4 failed"
    assert s.longestCommonPrefix(["abab","aba",""]) == "", "Test case 5 failed"
    
    print("All test cases passed")
