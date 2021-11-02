#
# @lc app=leetcode.cn id=21 lang=python3
#
# [21] 合并两个有序链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        # 链表上递归
        # 如果其中一个头节点对应的 val 值为 null，那么返回另外一个
        if l1 is None:
            return l2
        elif l2 is None:
            return l1
        # 都不为空，那么选择头节点对应值更小的作为返回的最终链表的头节点，就地修改，递归调用本函数，
        elif l1.val < l2.val:
            l1.next = self.mergeTwoLists(l1.next, l2)
            return l1
        else:
            l2.next = self.mergeTwoLists(l1, l2.next)
            return l2
# @lc code=end

