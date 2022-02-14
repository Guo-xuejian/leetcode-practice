#
# @lc app=leetcode.cn id=652 lang=python3
#
# [652] 寻找重复的子树
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def __init__(self):
        self.res = []
        self.all_tree = {}

    def findDuplicateSubtrees(self, root: Optional[TreeNode]) -> List[Optional[TreeNode]]:
        if root is None:
            return self.res
        # 深度优先，写入每一种树的序列化至 all_tree 并记录次数，1 或 2 次记录，2 次(重复)写入 res
        self.traverse(root)
        return self.res

    def traverse(self, root):
        # 序列化二叉树,采用后续遍历
        if root is None:    # 空节点，用 # 号代替
            return "#"

        left = self.traverse(root.left)
        right = self.traverse(root.right)

        # 描述整棵树
        subtree = left + "," + right + "," + str(root.val)

        if subtree in self.all_tree.keys() and self.all_tree[subtree] == 1:
            # 出现重复则写入 res，并且后续不会再次写入
            self.res.append(root)
            self.all_tree[subtree] += 1
        elif subtree not in self.all_tree.keys():
            # 则出现一次
            self.all_tree[subtree] = 1

        return subtree

# @lc code=end
