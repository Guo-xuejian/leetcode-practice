#
# @lc app=leetcode.cn id=1022 lang=python3
#
# [1022] 从根到叶的二进制数之和
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumRootToLeaf(self, root: Optional[TreeNode]) -> int:
        self.res = 0
        def dfs(node: TreeNode, curr: int):
            if node.left:
                dfs(node.left, curr * 2 + node.val)
            if node.right:
                dfs(node.right, curr * 2 + node.val)
            if not node.left and not node.right:
                self.res += curr * 2 + node.val
        if root.left:
            dfs(root.left, root.val)
        if root.right:
            dfs(root.right, root.val)
        if not root.left and not root.right:
            return root.val
        return self.res
# @lc code=end

