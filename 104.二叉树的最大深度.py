#
# @lc app=leetcode.cn id=104 lang=python3
#
# [104] 二叉树的最大深度
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        # 从两个分支中选出一个最大的，对于子节点也是如此，递归调用即可
        if not root:
            return 0
        else:
            left_length = self.maxDepth(root.left)
            right_length = self.maxDepth(root.right)
            return max(left_length, right_length) + 1   # 对于叶子节点需要额外加上
# @lc code=end

