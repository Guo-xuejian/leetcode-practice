#
# @lc app=leetcode.cn id=234 lang=python3
#
# [234] 回文链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        nodes = []
        while head:
            nodes.append(head.val)
            head = head.next
        left, right = 0, len(nodes) - 1
        while left < right:
            if nodes[left] == nodes[right]:
                left += 1
                right -= 1
            else:
                return False
        return True
# @lc code=end

