#
# @lc app=leetcode.cn id=998 lang=python3
#
# [998] 最大二叉树 II
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def insertIntoMaxTree(self, root: Optional[TreeNode], val: int) -> Optional[TreeNode]:
        parent, cur = None, root
        while cur:
            if val > cur.val:
                if not parent:
                    return TreeNode(val, root, None)
                node = TreeNode(val, cur, None)
                parent.right = node
                return root
            else:
                parent = cur
                cur = cur.right
        
        parent.right = TreeNode(val)
        return root
# @lc code=end

