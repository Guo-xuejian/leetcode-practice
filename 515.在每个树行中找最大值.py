#
# @lc app=leetcode.cn id=515 lang=python3
#
# [515] 在每个树行中找最大值
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def largestValues(self, root: Optional[TreeNode]) -> List[int]:
        # 标准的BFS
        if not root:
            return []
        queue = [root, ]
        res = []
        while len(queue) != 0:
            max_num = queue[0].val
            temp = []
            for node in queue:
                # 逐层遍历
                max_num = max_num if node.val <= max_num else node.val
                if node.left:
                    temp.append(node.left)
                if node.right:
                    temp.append(node.right)
            res.append(max_num)
            queue = temp[:]
        return res
# @lc code=end

