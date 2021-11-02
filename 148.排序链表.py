#
# @lc app=leetcode.cn id=148 lang=python3
#
# [148] 排序链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        temp1 = temp2 = head
        res = []        # 用一个列表 res 写入所有的元素，排序后依次覆盖原值
        while head:
            res.append(head.val)
            head = head.next
        res  = sorted(res)  # 排序
        for one in res:
            temp2.val = one
            temp2 = temp2.next
        return temp1
# @lc code=end

