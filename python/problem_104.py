from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        
        depth = self.dfs(root, 1)
        return depth

    def dfs(self, node, currentDepth):
        if node.left is None and node.right is None:
            return currentDepth
        
        leftDepth = currentDepth
        rightDepth = currentDepth

        if node.left is not None:
            leftDepth = self.dfs(node.left, currentDepth+1)

        if node.right is not None:
            rightDepth = self.dfs(node.right, currentDepth+1)

        if leftDepth > rightDepth:
            return leftDepth
        else:
            return rightDepth


if __name__ == "__main__":
    s = Solution()
    inputTree = TreeNode(3, TreeNode(9), TreeNode(
        20, TreeNode(15), TreeNode(7)))
    assert s.maxDepth(inputTree) == 3, "Test case 1 failed"

    inputTree = TreeNode(1, None, TreeNode(2))
    assert s.maxDepth(inputTree) == 2, "Test case 1 failed"

    print("All test cases passed")
