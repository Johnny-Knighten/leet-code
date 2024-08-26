from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if p is None and q is None:
            return True

        if p and q and p.val == q.val:
            return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)

        return False


if __name__ == "__main__":
    s = Solution()
    inputTree1 = TreeNode(1, TreeNode(2), TreeNode(3))
    inputTree2 = TreeNode(1, TreeNode(2), TreeNode(3))
    assert s.isSameTree(inputTree1, inputTree2), "Test case 1 failed"

    inputTree1 = TreeNode(1, TreeNode(2), None)
    inputTree2 = TreeNode(1, None, TreeNode(2))
    assert not s.isSameTree(inputTree1, inputTree2), "Test case 2 failed"

    inputTree1 = TreeNode(1, TreeNode(2), TreeNode(1))
    inputTree2 = TreeNode(1, TreeNode(1), TreeNode(2))
    assert not s.isSameTree(inputTree1, inputTree2), "Test case 3 failed"

    print("All test cases passed")
