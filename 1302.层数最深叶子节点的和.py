#
# @lc app=leetcode.cn id=1302 lang=python3
#
# [1302] 层数最深叶子节点的和
#
# https://leetcode.cn/problems/deepest-leaves-sum/description/
#
# algorithms
# Medium (82.73%)
# Likes:    87
# Dislikes: 0
# Total Accepted:    26.1K
# Total Submissions: 31.6K
# Testcase Example:  '[1,2,3,4,5,null,6,7,null,null,null,null,8]'
#
# 给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
# 
# 
# 
# 示例 1：
# 
# 
# 
# 
# 输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
# 输出：15
# 
# 
# 示例 2：
# 
# 
# 输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
# 输出：19
# 
# 
# 
# 
# 提示：
# 
# 
# 树中节点数目在范围 [1, 10^4] 之间。
# 1 
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
    def deepestLeavesSum(self, root: Optional[TreeNode]) -> int:
        q = deque([root])
        while q:
            ans = 0
            for _ in range(len(q)):
                node = q.popleft()
                ans += node.val
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
        return ans
# @lc code=end

