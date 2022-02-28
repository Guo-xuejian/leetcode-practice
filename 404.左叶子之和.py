#
# @lc app=leetcode.cn id=404 lang=python3
#
# [404] 左叶子之和
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumOfLeftLeaves(self, root: TreeNode) -> int:
        if not root:
            return 0
        res = 0

        if root.left and not root.left.left and not root.left.right:
            res += root.left.val
        
        return self.sumOfLeftLeaves(root.left) + self.sumOfLeftLeaves(root.right) + res
# @lc code=end

