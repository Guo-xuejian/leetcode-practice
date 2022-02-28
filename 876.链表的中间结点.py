#
# @lc app=leetcode.cn id=876 lang=python3
#
# [876] 链表的中间结点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        # dummy_node = head
        # node_count = 0
        # while head:
        #     node_count += 1
        #     head = head.next
        # node_count = node_count // 2
        # while node_count != 0:
        #     dummy_node = dummy_node.next
        #     node_count -= 1
        # return dummy_node
        # 快慢指针
        pre_node = head
        while head and head.next:
            pre_node = pre_node.next
            head = head.next.next
        return pre_node
# @lc code=end

