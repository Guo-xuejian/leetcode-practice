#
# @lc app=leetcode.cn id=83 lang=python3
#
# [83] 删除排序链表中的重复元素
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        dummy_node = head
        if not head:
            return head
        pre_node = head
        head = head.next
        while head:
            if head.val == pre_node.val:
                pre_node.next = head.next
            else:
                pre_node = head
            head = head.next
        return dummy_node
# @lc code=end
