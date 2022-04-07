#
# @lc app=leetcode.cn id=429 lang=python3
#
# [429] N 叉树的层序遍历
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
    def levelOrder(self, root: 'Node') -> List[List[int]]:
        res = []
        if not root:
            return res
        queue = [root, ]
        while len(queue) > 0:
            temp = queue[:]
            queue = []
            curr = []
            for node in temp:
                curr.append(node.val)
                for next_layer_node in node.children:
                    queue.append(next_layer_node)
            res.append(curr)
        return res

# @lc code=end

