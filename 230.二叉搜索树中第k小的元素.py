#
# @lc app=leetcode.cn id=230 lang=python3
#
# [230] 二叉搜索树中第K小的元素
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        # 写入列表排序，获取对应索引即可
        res = []
        def insertIntoRes(root):
            if root:    # 节点不为空才会写入
                res.append(int(root.val))
                insertIntoRes(root.left)
                insertIntoRes(root.right)
        insertIntoRes(root)
        # 排序后获取 k-1 元素，因为题目要求从 1 开始计数
        return sorted(res)[k-1]
# @lc code=end

