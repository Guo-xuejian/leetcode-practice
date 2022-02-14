# 剑指 Offer 33. 二叉搜索树的后序遍历序列
# 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

 

# 参考以下这颗二叉搜索树：

#      5
#     / \
#    2   6
#   / \
#  1   3
# 示例 1：

# 输入: [1,6,3,2,5]
# 输出: false
# 示例 2：

# 输入: [1,3,2,6,5]
# 输出: true
 

# 提示：

# 数组长度 <= 1000


class Solution:
    def verifyPostorder(self, postorder: List[int]) -> bool:
        # 一直弹出一个根节点(实际有没有子节点不重要，反正可以是个根节点)
        # 遵循 根 右 左 的镜像顺序，递归时修改上下限(根节点是左子树的上限以及右子树的下限)
        # 如果存在一个不满足上下限关系，则退出，最后 postorder 不为空就证明不是
        def build(postorder: List[int], up: int, down: int):
            if not postorder:
                return
            curr_val = postorder[-1]
            if not down < curr_val < up:
                return
            postorder.pop() # 根
            build(postorder, up, curr_val) # 右
            build(postorder, curr_val, down) # 左

        build(postorder, float("inf"), -float("inf"))
        return not postorder