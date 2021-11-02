#
# @lc app=leetcode.cn id=437 lang=python3
#
# [437] 路径总和 III
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: TreeNode, targetSum: int) -> int:
        number_list = []
        def down(root, number_list):
            if not root:
                return 0
            # 存储这个路径上所有数字相加的结果，如果其中有等于targetNum的，那就是满足条件的
            # 同时要保证存储的结果中每一个都是一条路径下的综合结果
            number_list = [number + root.val for number in number_list]
            # 保存该节点
            number_list.append(root.val)
            count = 0
            for number in number_list:
                if number == targetSum:
                    count += 1
            # 递归调用，二叉树都是这样子递归的，一个套一个
            return count + down(root.left, number_list) + down(root.right, number_list)
        return down(root, number_list)
# @lc code=end

