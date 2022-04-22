# 剑指 Offer 25. 合并两个排序的链表
# 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

# 示例1：

# 输入：1->2->4, 1->3->4
# 输出：1->1->2->3->4->4
# 限制：

# 0 <= 链表长度 <= 1000


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:
        if not l1 and not l2:
            return l1
        dummy_node = ListNode(0)
        pre = dummy_node
        while l1 or l2:
            if l1 and l2:
                if l1.val > l2.val:
                    pre.val = l2.val
                    l2 = l2.next
                else:
                    pre.val = l1.val
                    l1 = l1.next
                temp = ListNode(0)
                pre.next = temp
                pre = temp
            elif not l1:
                pre.val = l2.val
                l2 = l2.next
                pre.next = l2
                break
            else:
                pre.val = l1.val
                l1 = l1.next
                pre.next = l1
                break
        return dummy_node