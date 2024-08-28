class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        sCharLastIndexMap = dict()
        tCharLastIndexMap = dict()
        for i in range(len(s)):
            if sCharLastIndexMap.get(s[i], 0) != tCharLastIndexMap.get(t[i], 0):
                return False

            sCharLastIndexMap[s[i]] = sCharLastIndexMap.get(s[i], i) + 1
            tCharLastIndexMap[t[i]] = tCharLastIndexMap.get(t[i], i) + 1

        return True


if __name__ == "__main__":
    s = Solution()
    input1 = "egg"
    input2 = "add"
    assert s.isIsomorphic(input1, input2), "Test case 1 failed"

    input1 = "foo"
    input2 = "bar"
    assert not s.isIsomorphic(input1, input2), "Test case 2 failed"

    input1 = "paper"
    input2 = "title"
    assert s.isIsomorphic(input1, input2), "Test case 3 failed"

    input1 = "bbbaaaba"
    input2 = "aaabbbba"
    assert not s.isIsomorphic(input1, input2), "Test case 4 failed"

    print("All test cases passed")
