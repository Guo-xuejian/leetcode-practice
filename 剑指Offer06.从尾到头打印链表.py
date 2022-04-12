# 剑指 Offer 06. 从尾到头打印链表
# 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。


# 示例 1：

# 输入：head = [1,3,2]
# 输出：[2,3,1]


# 限制：

# 0 <= 链表长度 <= 10000


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reversePrint(self, head: ListNode) -> List[int]:
        res = []
        while head:
            res.append(head.val)
            head = head.next
        # return res[::-1]
        length = len(res)
        if length == 0:
            return res
        index = 0
        while index <= length // 2 - 1:
            res[index], res[length - 1 -
                            index] = res[length - 1 - index], res[index]
            index += 1
        return res

# 2022-04-13
# class Solution:
#     def reversePrint(self, head: ListNode) -> List[int]:
#         res = []
#         while head:
#             res.append(head.val)
#             head = head.next
#         return res[::-1]