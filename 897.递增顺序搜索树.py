#
# @lc app=leetcode.cn id=897 lang=python3
#
# [897] 递增顺序搜索树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def increasingBST(self, root: TreeNode) -> TreeNode:
        all_val = []

        def inOrder(root):
            if root:
                inOrder(root.left)
                all_val.append(root.val)
                inOrder(root.right)

        inOrder(root)

        pre_head = TreeNode()
        curr_node = pre_head

        for val in all_val:
            curr_node.right = TreeNode(val=val)
            curr_node = curr_node.right

        return pre_head.right

# @lc code=end
