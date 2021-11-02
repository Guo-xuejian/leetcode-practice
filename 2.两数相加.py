#
# @lc app=leetcode.cn id=2 lang=python3
#
# [2] 两数相加
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        # 两位相加
        # 创建一个链表（最后是NoneNone）
        dummy = p = ListNode(None)
        Sum = 0
        # Sum 不是0的时候进行loop
        while l1 or l2 or Sum != 0:
            # 更新Sum
            Sum += (l1.val if l1 else 0) + (l2.val if l2 else 0)
            # 给链表p.next赋值，值为Sum%10
            p.next = ListNode(Sum % 10)
            # 移动p的位置，是指针移动到下一个位置
            p = p.next
            # 如果l1非空，移动到下一个位置
            if l1: l1 = l1.next
            if l2: l2 = l2.next
            # Sum更新，取整,Sum的值留到高位相加，Sum也是最终结束的判定条件
            Sum = Sum//10

        return dummy.next
# @lc code=end

