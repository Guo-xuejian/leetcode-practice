#
# @lc app=leetcode.cn id=1325 lang=python3
#
# [1325] 删除给定值的叶子节点
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def removeLeafNodes(self, root: Optional[TreeNode], target: int) -> Optional[TreeNode]:
        if not root:
            return None
        root.left = self.removeLeafNodes(root.left, target)     # 递归删减左子节点
        root.right = self.removeLeafNodes(root.right, target)   # 递归删减右子节点

        # 删除至父节点成为叶子节点，那么父节点也应该被删除
        if not root.left and not root.right and root.val == target:
            return None
        # 但凡一个条件不满足就正常返回
        return root
# @lc code=end
