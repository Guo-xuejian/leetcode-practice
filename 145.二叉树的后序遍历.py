#
# @lc app=leetcode.cn id=145 lang=python3
#
# [145] 二叉树的后序遍历
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        res = []
        def postOrder(node):
            if node != None:    # 节点不为空，顺序为左子树、右子树、根节点
                postOrder(node.left)
                postOrder(node.right)
                res.append(node.val)
        
        postOrder(root)
        return res
# @lc code=end

