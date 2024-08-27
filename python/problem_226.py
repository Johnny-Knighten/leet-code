from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root or (not root.left and not root.right):
            return root

        root.left, root.right = root.right, root.left
        root.left = self.invertTree(root.left)
        root.right = self.invertTree(root.right)
        return root


if __name__ == "__main__":
    s = Solution()
    input = TreeNode(4, TreeNode(2, TreeNode(1), TreeNode(3)),
                     TreeNode(7, TreeNode(6),  TreeNode(9)))
    expected = TreeNode(4, TreeNode(7, TreeNode(9), TreeNode(6)),
                        TreeNode(2, TreeNode(3),  TreeNode(1)))
    assert s.invertTree(input) == expected, "Test case 1 failed"

    input = TreeNode(2, TreeNode(1), TreeNode(3))
    expected = TreeNode(2, TreeNode(3), TreeNode(1))
    assert s.invertTree(input) == expected, "Test case 2 failed"
