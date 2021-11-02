#
# @lc app=leetcode.cn id=144 lang=python3
#
# [144] 二叉树的前序遍历
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def preorderTraversal(self, root: TreeNode) -> List[int]:
        res = []
        # 二叉树的遍历在函数内部实现方法去递归
        def preOrder(node):
            if node != None:        # 只要节点不为空就写入 
                res.append(node.val)
                preOrder(node.left)
                preOrder(node.right)
        
        preOrder(root)
        return res
# @lc code=end

