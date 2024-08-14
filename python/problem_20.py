class Solution:
    def isValid(self, s: str) -> bool:
        paras = []
        for char in s:
            if char in ["{", "(", "["]:
                paras.append(char)
            else:
                if len(paras) == 0:
                    return False
                if paras[-1] != "{" and char == "}":
                    return False
                if paras[-1] != "[" and char == "]":
                    return False
                if paras[-1] != "(" and char == ")":
                    return False
                paras.pop()

        return len(paras) == 0
                

if __name__ == "__main__":
    s = Solution()
    assert s.isValid("()") == True, "Test case 1 failed"
    assert s.isValid("()[]{}") == True, "Test case 2 failed"
    assert s.isValid("(]") == False, "Test case 3 failed"
    assert s.isValid("((") == False, "Test case 4 failed"

    print("All test cases passed")
