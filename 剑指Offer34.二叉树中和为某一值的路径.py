# 剑指 Offer 34. 二叉树中和为某一值的路径
# 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

# 叶子节点 是指没有子节点的节点。

 

# 示例 1：



# 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
# 输出：[[5,4,11,2],[5,8,4,5]]
# 示例 2：



# 输入：root = [1,2,3], targetSum = 5
# 输出：[]
# 示例 3：

# 输入：root = [1,2], targetSum = 0
# 输出：[]
 

# 提示：

# 树中节点总数在范围 [0, 5000] 内
# -1000 <= Node.val <= 1000
# -1000 <= targetSum <= 1000


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: TreeNode, target: int) -> List[List[int]]:
        ret = list()
        path = list()
        
        def dfs(root: TreeNode, target: int):
            if not root:
                return
            path.append(root.val)
            target -= root.val
            if not root.left and not root.right and target == 0:
                ret.append(path[:])
            dfs(root.left, target)
            dfs(root.right, target)
            path.pop()
        
        dfs(root, target)
        return ret