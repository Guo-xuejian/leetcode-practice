#
# @lc app=leetcode.cn id=19 lang=python3
#
# [19] 删除链表的倒数第 N 个结点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        layer = 0
        temp1 = temp2 = ListNode(0,head)
        # 先遍历拿到长度计算出删除的位置
        while head:
            layer += 1
            head = head.next
        #delete_index = layer - n 计算出需要删除的元素索引
        for _ in range(0,layer-n):    # 但是删除该元素需要到它的前一个元素，修改它的前一个元素的 next 指向即可
            temp1 = temp1.next
        # 执行 next 指向，直接跳过删除元素即可
        temp1.next = temp1.next.next
        return temp2.next
# @lc code=end

