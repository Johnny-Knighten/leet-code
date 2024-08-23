class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        magazineDict = dict()
        for letter in magazine:
            magazineDict[letter] = magazineDict.get(letter, 0) + 1

        ransomNoteDict = dict()
        for letter in ransomNote:
            ransomNoteDict[letter] = ransomNoteDict.get(letter, 0) + 1

        for key, val in ransomNoteDict.items():
            if key not in magazineDict or magazineDict[key] < val:
                return False

        return True


if __name__ == "__main__":
    s = Solution()
    assert not s.canConstruct("a", "b"), "Test case 1 failed"
    assert not s.canConstruct("aa", "ab"), "Test case 2 failed"
    assert s.canConstruct("aa", "aab"), "Test case 3 failed"
    print("All test cases passed")
