#
# @lc app=leetcode.cn id=543 lang=python3
#
# [543] 二叉树的直径
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        self.res = 1
        def depth(node):
            # 寻找每个子节点的最大长度，加总即可
            if not node:
                return 0
            L, R = depth(node.left), depth(node.right)
            if R + L + 1 > self.res:
                self.res = R + L + 1
            return max(L, R) + 1
        depth(root)
        return self.res - 1
# @lc code=end

