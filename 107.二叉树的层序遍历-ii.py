#
# @lc app=leetcode.cn id=107 lang=python3
#
# [107] 二叉树的层序遍历 II
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrderBottom(self, root: TreeNode) -> List[List[int]]:
        res = []

        def dfs(root, depth):
            if root == None:
                return []
            # 后续的递归可能导致 len(res) != depth，做一下判定，确定是否需要加上下一层节点的列表
            if len(res) == depth:
                res.append([])
            res[depth].append(root.val)
            # 左子节点，进入下一层
            if root.left:
                dfs(root.left, depth + 1)
            # 同上，右子节点
            if root.right:
                dfs(root.right, depth + 1)
        dfs(root, 0)     # 从根节点开始，层数为 0 
        
        # 需要翻转
        return res[::-1]
# @lc code=end

