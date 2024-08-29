from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True

        return self.recurIsSymmetric(root.left, root.right)

    def recurIsSymmetric(self, node1: Optional[TreeNode], node2: Optional[TreeNode]) -> bool:
        if node1 is None and node2 is None:
            return True

        if node1 is None or node2 is None or node1.val != node2.val:
            return False

        return self.recurIsSymmetric(node1.left, node2.right) and self.recurIsSymmetric(node1.right, node2.left)


if __name__ == "__main__":
    s = Solution()
    inputTree1 = TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(
        4)), TreeNode(2, TreeNode(4), TreeNode(3)))
    assert s.isSymmetric(inputTree1), "Test case 1 failed"

    inputTree1 = TreeNode(1, TreeNode(2, right=TreeNode(3)),
                          TreeNode(2, right=TreeNode(3)))
    assert not s.isSymmetric(inputTree1), "Test case 2 failed"

    print("All test cases passed")
