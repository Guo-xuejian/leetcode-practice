#
# @lc app=leetcode.cn id=94 lang=python3
#
# [94] 二叉树的中序遍历
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        final_list = []
        def inOrder(node):
            if not node:
                return
            inOrder(node.left)
            final_list.append(node.val)
            inOrder(node.right)
        inOrder(root)
        return final_list
# @lc code=end

