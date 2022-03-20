#
# @lc app=leetcode.cn id=653 lang=python3
#
# [653] 两数之和 IV - 输入 BST
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        # 一样的哈希表
        num_dict = {}
        self.success = False
        def writeToDict(node):
            # 空或者已经找到则不继续
            if not node or self.success:
                return
            if node.val in num_dict:
                self.success = True
                return
            else:
                num_dict[k - node.val] = 1
            writeToDict(node.left)
            writeToDict(node.right)
        writeToDict(root)
        return self.success
# @lc code=end

