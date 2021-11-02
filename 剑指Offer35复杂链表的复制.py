"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""
class Solution:
    def copyRandomList(self, head: 'Node') -> 'Node':
        if not head:
            return None
        
        p = head
        dic = dict()
        # 将每一个节点都放到字典中去，实际上并不是所有的内容都放进去，只是节点的值创建并放入，next节点和random节点都在下面以访问原链表的方式创建
        while p:
            # 创建节点
            dic[p] = Node(p.val)
            p = p.next
        
        p = head
        while p:
            # 因为是复制，不是生成，所以结构上要保持一致，无论是next节点还是random节点
            dic[p].next = dic.get(p.next)
            dic[p].random = dic.get(p.random)
            p = p.next
        
        return dic[head]
