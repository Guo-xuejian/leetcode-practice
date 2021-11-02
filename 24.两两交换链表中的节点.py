#
# @lc app=leetcode.cn id=24 lang=python3
#
# [24] 两两交换链表中的节点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        dummyHead = ListNode(0)
        dummyHead.next = head
        temp = dummyHead
        while temp.next and temp.next.next:
            node1 = temp.next   # 一号节点
            node2 = temp.next.next  # 二号节点
            temp.next = node2   # 初始节点指向二号节点
            node1.next = node2.next     # 一号节点指向二号节点的下一节点
            node2.next = node1  # 二号节点指向一号节点 完成交换
            # 由于是两两交换，也就是一二交换，三四交换，二三不叫唤，所以需要进行节点的跳跃，不能每个节点都来一次
            temp = node1    # 由于node后移一位，所以实际上跳过了一个不应该被执行的节点，有意思！！
        return dummyHead.next

        
            
# @lc code=end

