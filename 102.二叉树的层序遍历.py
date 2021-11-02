#
# @lc app=leetcode.cn id=102 lang=python3
#
# [102] 二叉树的层序遍历
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if root == None:
            return []
        res = []
        q = [root]  # 第一层有一个根节点
        index = 0
        # 难点在于如何记住每一层======> 把每一层都写入队列，遍历时将其子树写入新队列后覆盖
        while len(q) > 0:
            p = []      # 即将替代 p 的切片
            res.append([])
            for j in range(len(q)):     # 根据当前层所拥有的节点数量
                node = q[j]
                res[index].append(node.val)
                if node.left != None:   # 写入左子树
                    p.append(node.left)
                if node.right != None:  # 写入右子树
                    p.append(node.right)
            # 迭代至下一层
            q = p
            index += 1
        return res
# @lc code=end

