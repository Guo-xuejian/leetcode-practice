# 92. 反转链表 II
# 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
 

# 示例 1：


# 输入：head = [1,2,3,4,5], left = 2, right = 4
# 输出：[1,4,3,2,5]
# 示例 2：

# 输入：head = [5], left = 1, right = 1
# 输出：[5]
 

# 提示：

# 链表中节点数目为 n
# 1 <= n <= 500
# -500 <= Node.val <= 500
# 1 <= left <= right <= n
 

# 进阶： 你可以使用一趟扫描完成反转吗？


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: ListNode, left: int, right: int) -> ListNode:
        if not head:
            return head
        dummy_node = ListNode(0)
        dummy_node.next = head
        temp = dummy_node
        for _ in range(left - 1):
            temp = temp.next
            head = head.next
        start_pre = temp
        start = head
        temp = head.next
        pre = head
        for _ in range(right - left):
            head = temp
            temp = temp.next
            head.next = pre
            pre = head
        start_pre.next = pre
        start.next = temp
        return dummy_node.next
