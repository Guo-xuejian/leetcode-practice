#
# @lc app=leetcode.cn id=1161 lang=python3
#
# [1161] 最大层内元素和
#
# https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/description/
#
# algorithms
# Medium (62.86%)
# Likes:    59
# Dislikes: 0
# Total Accepted:    14.4K
# Total Submissions: 22.6K
# Testcase Example:  '[1,7,0,7,-8,null,null]'
#
# 给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
# 
# 请返回层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。
# 
# 
# 
# 示例 1：
# 
# 
# 
# 
# 输入：root = [1,7,0,7,-8,null,null]
# 输出：2
# 解释：
# 第 1 层各元素之和为 1，
# 第 2 层各元素之和为 7 + 0 = 7，
# 第 3 层各元素之和为 7 + -8 = -1，
# 所以我们返回第 2 层的层号，它的层内元素之和最大。
# 
# 
# 示例 2：
# 
# 
# 输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
# 输出：2
# 
# 
# 
# 
# 提示：
# 
# 
# 树中的节点数在 [1, 10^4]范围内
# -10^5 <= Node.val <= 10^5
# 
# 
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxLevelSum(self, root: Optional[TreeNode]) -> int:
        # bfs
        node_stack = [root, ]
        max_num = -float("inf")
        res = 1
        layer = 1
        while len(node_stack) != 0:
            curr_num = sum([node.val for node in node_stack])
            if curr_num > max_num:
                res = layer
            max_num = max(max_num, curr_num)
            temp = []
            for node in node_stack:
                if node.left:
                    temp.append(node.left)
                if node.right:
                    temp.append(node.right)
            node_stack = temp[:]
            layer += 1
        return res

# 执行用时：252 ms, 在所有 Python3 提交中击败了88.89%的用户
# 内存消耗：19.6 MB, 在所有 Python3 提交中击败了69.93%的用户
# @lc code=end

