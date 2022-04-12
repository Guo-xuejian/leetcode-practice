#
# @lc app=leetcode.cn id=206 lang=python3
#
# [206] 反转链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        dummy_node = None
        while head:
            temp = head.next
            head.next = dummy_node
            dummy_node = head
            head = temp
        return dummy_node
# @lc code=end

