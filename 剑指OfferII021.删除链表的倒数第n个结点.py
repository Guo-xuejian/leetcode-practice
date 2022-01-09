# 剑指 Offer II 021. 删除链表的倒数第 n 个结点
# 给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

# 示例 1：

# 输入：head = [1,2,3,4,5], n = 2
# 输出：[1,2,3,5]
# 示例 2：

# 输入：head = [1], n = 1
# 输出：[]
# 示例 3：

# 输入：head = [1,2], n = 1
# 输出：[1]

# 提示：

# 链表中结点的数目为 sz
# 1 <= sz <= 30
# 0 <= Node.val <= 100
# 1 <= n <= sz

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        left = right = head
        count = 0
        # 维持两个指针的距离为 n ，那么当right指向链表最后的元素时，left也就指向了当属第二个，left.next = left.next.next完事
        while count < n:    # right 先走距离 n，距离差出现
            right = right.next
            count += 1
        if not right:   # 超出最大长度，也就是 n==长度，删除的也就是 head，返回 head.next
            return head.next
        while right.next:   # left right 同步移动，维持距离差 n
            left = left.next
            right = right.next
        # left 指向倒数第 n-1 个元素，覆盖第 n 个
        left.next = left.next.next
        return head
