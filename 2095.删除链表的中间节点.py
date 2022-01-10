#
# @lc app=leetcode.cn id=2095 lang=python3
#
# [2095] 删除链表的中间节点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head.next:   # 单节点只能删除该节点，最后为 None
            return None

        # 快慢指针，由于中间节点，而slow和fast一起走的话只能满足slow到中间节点，需要一个pre_slow来绑定slow前一个位置，最终也就是中间节点的前一个节点，
        slow, fast, pre_slow = head, head, None
        while fast and fast.next:
            fast = fast.next.next   # 一次走两个节点，倍数差 2
            pre_slow = slow     # 记录当前中间节点的前一个节点
            slow = slow.next    # 一次一个节点，最终就是中间节点
        pre_slow.next = pre_slow.next.next
        return head     # 返回头节点

# @lc code=end

