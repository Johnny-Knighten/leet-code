from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if root is None:
            return False

        if root.left is None and root.right is None and targetSum-root.val == 0:
            return True

        return self.hasPathSum(root.left, targetSum-root.val) or self.hasPathSum(root.right, targetSum-root.val)


if __name__ == "__main__":
    s = Solution()

    input = TreeNode(5, TreeNode(4, TreeNode(11, TreeNode(7), TreeNode(
        2))), TreeNode(8, TreeNode(13), TreeNode(4, None, TreeNode(1))))
    targetSum = 22
    assert s.hasPathSum(input, targetSum), "Test case 1 failed"

    input = TreeNode(1, TreeNode(2), TreeNode(3))
    targetSum = 5
    assert not s.hasPathSum(input, targetSum), "Test case 2 failed"

    input = None
    targetSum = 0
    assert not s.hasPathSum(input, targetSum), "Test case 3 failed"

    print("All test cases passed")
