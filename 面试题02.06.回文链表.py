# 面试题 02.06. 回文链表
# 编写一个函数，检查输入的链表是否是回文的。

 

# 示例 1：

# 输入： 1->2
# 输出： false 
# 示例 2：

# 输入： 1->2->2->1
# 输出： true 
 

# 进阶：
# 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        arr = []

        while head:
            arr.append(head.val)
            head = head.next
        
        return arr == arr[::-1]