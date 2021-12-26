#
# @lc app=leetcode.cn id=1609 lang=python3
#
# [1609] 奇偶树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isEvenOddTree(self, root: Optional[TreeNode]) -> bool:
        queue = [root]
        level = 0
        while queue:
            # 奇数层递减给出较大值，偶数层递增给出较小值0（题目给出范围大于等于 1）
            prev = float('inf') if level % 2 else 0
            next = []
            for node in queue:
                val = node.val
                # 偶数层不出现偶数，奇数层不出现奇数
                if val % 2 == level % 2:
                    return False
                # 偶数层单调递增，递减则不是奇偶树
                if level % 2 == 0 and val <= prev:
                    return False
                # 奇数层单调递减，递增则不符合要求
                if level % 2 == 1 and val >= prev:
                    return False
                prev = val
                if node.left:
                    next.append(node.left)
                if node.right:
                    next.append(node.right)
            queue = next    # 覆盖当前层，叶子节点层时 next 为空则退出循环
            level += 1      # 下一层
        return True


# @lc code=end
