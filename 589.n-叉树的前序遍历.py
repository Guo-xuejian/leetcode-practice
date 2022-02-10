#
# @lc app=leetcode.cn id=589 lang=python3
#
# [589] N 叉树的前序遍历
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
    # 递归实现
    # def preorder(self, root: 'Node') -> List[int]:
    #     res = []
    #     def dfs(root):
    #         if root:
    #             res.append(root.val)
    #             for node in root.children:
    #                 dfs(node)
    #     dfs(root)
    #     return res
        
    def preorder(self, root: 'Node') -> List[int]:
        if root is None:
            return []
        stack, res = [root,], []
        while stack:
            root = stack.pop()
            res.append(root.val)
            stack.extend(root.children[::-1])   # 逆序写入，保证下次一 pop 出的元素在栈顶
        return res
# @lc code=end

