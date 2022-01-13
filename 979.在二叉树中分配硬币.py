#
# @lc app=leetcode.cn id=979 lang=python3
#
# [979] 在二叉树中分配硬币
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def distributeCoins(self, root):
        # 取决于每个节点的子节点需不需要，大于 1 需要移动，为 0 也需要（从父节点获取）
        self.result = 0

        def dfs(root):
            if not root:
                return 0
            left, right = dfs(root.left), dfs(root.right)
            self.result += abs(left) + abs(right)
            # -1 是为了保证目前节点至少有一个硬币
            return root.val + left + right - 1

        dfs(root)
        return self.result
# @lc code=end
