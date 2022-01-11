#
# @lc app=leetcode.cn id=606 lang=python3
#
# [606] 根据二叉树创建字符串
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def tree2str(self, root: Optional[TreeNode]) -> str:
        if not root:
            return ""
        if not root.left and not root.right:    # 无子节点
            return str(root.val)
        if not root.right:  # 左子节点可能存在，判断右子节点，因为该情况下如果右子节点不存在，那么右子节点对应括号缺省
            return str(root.val) + "(" + self.tree2str(root.left) + ")"
        # 都存在则递归写入
        return str(root.val) + "(" + self.tree2str(root.left) + ")(" + self.tree2str(root.right) + ")"
# @lc code=end
