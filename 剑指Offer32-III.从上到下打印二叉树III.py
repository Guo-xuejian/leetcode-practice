# 剑指 Offer 32 - III. 从上到下打印二叉树 III
# 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

 

# 例如:
# 给定二叉树: [3,9,20,null,null,15,7],

#     3
#    / \
#   9  20
#     /  \
#    15   7
# 返回其层次遍历结果：

# [
#   [3],
#   [20,9],
#   [15,7]
# ]
 

# 提示：

# 节点总数 <= 1000



# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        # 层序遍历
        res = []
        if not root:
            return res
        queue = [root, ]
        flag = True
        while len(queue) > 0:
            temp = queue[:]
            queue = []
            curr = []
            if flag:
                flag = False
                for node in temp:
                    curr.append(node.val)
                    if node.left:
                        queue.append(node.left)
                    if node.right:
                        queue.append(node.right)
            else:
                flag = True
                for node in temp[::-1]:
                    curr.append(node.val)
                    # 后续需要取反，所以先右后左
                    if node.right:
                        queue.append(node.right)
                    if node.left:
                        queue.append(node.left)
                # 由于反向遍历，导致 queue 顺序也产生变化，取反
                queue = queue[::-1]
            res.append(curr)
        return res