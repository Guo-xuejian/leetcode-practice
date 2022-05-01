#
# @lc app=leetcode.cn id=1305 lang=python3
#
# [1305] 两棵二叉搜索树中的所有元素
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def getAllElements(self, root1: TreeNode, root2: TreeNode) -> List[int]:
        def dfs(root, res):
            if root:
                dfs(root.left, res)
                res.append(root.val)
                dfs(root.right, res)
        
        ans, nums1, nums2 = [], [], []
        dfs(root1, nums1), dfs(root2, nums2)
        idx1 = idx2 = 0
        while idx1 < len(nums1) and idx2 < len(nums2):
            if nums1[idx1] <= nums2[idx2]:
                ans.append(nums1[idx1])
                idx1 += 1
            else:
                ans.append(nums2[idx2])
                idx2 += 1
        ans.extend(nums1[idx1:] + nums2[idx2:])
        return ans
        
# @lc code=end

