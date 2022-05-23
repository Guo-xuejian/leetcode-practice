#
# @lc app=leetcode.cn id=965 lang=python3
#
# [965] 单值二叉树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isUnivalTree(self, root: TreeNode) -> bool:
        pre_num = root.val

        def dfs(node: TreeNode) -> bool:
            if not node:    # 空则正常
                return False
            if node.val != pre_num:     # 不相等直接退出
                return True
            # 目前为止满足单值二叉树定义
            return dfs(node.left) or dfs(node.right)
        if dfs(root.left) or dfs(root.right):
            # 不满足条件
            return False
        return True
# @lc code=end

