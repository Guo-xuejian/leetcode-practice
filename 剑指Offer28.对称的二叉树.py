# 剑指 Offer 28. 对称的二叉树
# 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

# 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

#     1
#    / \
#   2   2
#  / \ / \
# 3  4 4  3
# 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

#     1
#    / \
#   2   2
#    \   \
#    3    3

 

# 示例 1：

# 输入：root = [1,2,2,3,4,4,3]
# 输出：true
# 示例 2：

# 输入：root = [1,2,2,null,3,null,3]
# 输出：false
 

# 限制：

# 0 <= 节点个数 <= 1000



# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isSymmetric(self, root: TreeNode) -> bool:
        def recur(A, B):
            if not A and not B:
                return True
            elif not A or not B:
                return False
            return A.val == B.val and recur(A.left, B.right) and recur(A.right, B.left)
        return recur(root, root)