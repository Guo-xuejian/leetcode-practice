#
# @lc app=leetcode.cn id=237 lang=python3
#
# [237] 删除链表中的节点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def deleteNode(self, node):
        """
        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """
        # 由于题目中说明是一个有效节点且不失末尾节点，那么就不需要判定Next为空
	    # 覆盖当前的节点为下一个节点的值，贪吃蛇向前走一步
        while node.next.next != None:   # 始终保持指针在最后一个节点之前的那一个节点
            node.val = node.next.val    # 覆盖当前节点值为下一个节点值
            node = node.next    # 指针移动，于是这个节点就被删除了
        # 但是最后一个节点被跳过了，需要覆盖其 next 指针为 None
        node.val = node.next.val
        node.next = None
        
# @lc code=end

