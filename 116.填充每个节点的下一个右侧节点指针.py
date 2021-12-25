#
# @lc app=leetcode.cn id=116 lang=python3
#
# [116] 填充每个节点的下一个右侧节点指针
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""
import collections


class Solution:
    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        if not root:
            return root
        # 初始化队列同时将第一层节点加入队列中，即根节点
        Q = collections.deque([root])
        # 外层的 while 循环迭代层数
        while Q:
            # 记录当前队列的大小
            size = len(Q)
            # 遍历目前层的所有节点
            for i in range(size):
                # 从队首取出元素
                node = Q.popleft()  # 每一次取出来的就是相对左子节点，Q[0]也就是题目要求的右子节点
                # 连接
                if i < size - 1:
                    node.next = Q[0]
                # 拓展下一层节点
                if node.left:
                    Q.append(node.left)
                if node.right:
                    Q.append(node.right)
        # 返回根节点即可
        return root
# @lc code=end
