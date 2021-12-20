# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def inorderSuccessor(self, root: 'TreeNode', p: 'TreeNode') -> 'TreeNode':
        res = None

        while root:
            if root.val > p.val:    # 大走左节点，因为二叉搜索树左子树都小于根节点
                res = root
                root = root.left    # 直到找到小于对应根节点的值中最小的
            else:                   # 否则走右子树
                root = root.right
        return res