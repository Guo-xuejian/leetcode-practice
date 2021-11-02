# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# 迭代 双指针
class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        cur, pre = head, None
        while cur:
            # 先保存下一个节点，因为马上就要修改它的指向，我们直接就是把next指向改为反向
            tmp = cur.next
            # 修改指向，第一次指向Null也就是最后的结尾
            cur.next = pre
            # 保存下一个元素的指向为当前元素,到最后也就是头节点
            pre = cur
            # 下一个节点
            cur = tmp
        return pre


# 递归 两者思想其实都是相似的，遍历的同时修改节点指向
class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        def recur(cur, pre):
            if not cur: return pre     # 终止条件
            res = recur(cur.next, cur) # 递归后继节点
            cur.next = pre             # 修改节点引用指向
            return res                 # 返回反转链表的头节点
        
        return recur(head, None)   

