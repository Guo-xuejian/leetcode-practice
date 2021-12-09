#
# @lc app=leetcode.cn id=559 lang=python3
#
# [559] N 叉树的最大深度
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution:
    def maxDepth(self, root: 'Node') -> int:
        # 一个空节点意味着上一节点为叶子节点，实际上是虚构的一个节点
        if not root:
            return 0
        max_child_depth = 0
        for child in root.children:
            child_depth = self.maxDepth(child)  # 递归
            if child_depth > max_child_depth:
                # 取子节点中深度较大者
                max_child_depth = child_depth
        # maxChildDepth 对应子节点大的最大深度，加上本身节点的高度 1
        return max_child_depth + 1
# @lc code=end
