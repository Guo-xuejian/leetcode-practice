# 剑指 Offer 55 - II. 平衡二叉树
# 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

 

# 示例 1:

# 给定二叉树 [3,9,20,null,null,15,7]

#     3
#    / \
#   9  20
#     /  \
#    15   7
# 返回 true 。

# 示例 2:

# 给定二叉树 [1,2,2,3,3,null,null,4,4]

#        1
#       / \
#      2   2
#     / \
#    3   3
#   / \
#  4   4
# 返回 false 。

 

# 限制：

# 0 <= 树的结点个数 <= 10000



# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isBalanced(self, root: TreeNode) -> bool:
        def height(root: TreeNode) -> int:
            if not root:
                return 0
            return max(height(root.left), height(root.right)) + 1

        if not root:
            return True
        return abs(height(root.left) - height(root.right)) <= 1 and self.isBalanced(root.left) and self.isBalanced(root.right)

# 题目：有效字符串匹配判断

# 给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：  任何左括号 ( 必须有相应的右括号 )。 任何右括号 ) 必须有相应的左括号 ( 。 左括号 ( 必须在对应的右括号之前 )。 * 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。 一个空字符串也被视为有效字符串。
def isValid(s: str) -> bool:
    letter_list = []
    for letter in s:
        if not letter_list:
            if letter == ")":
                # 前置无匹配左括号
                return False
            letter_list.append(letter)
        elif letter == ")":
            # 右括号匹配最近的左括号
            if letter_list[-1] == "*" or letter_list[-1] == "(":
                letter_list.pop()
            else:
                return False
        else:
            # 左括号入栈
            letter_list.append(letter)
    if letter_list[-1] == "(":
        return False
    return True


string = input()
print(string)
print(isValid(string))
# 1 名称错误 2 路径错误
# rm -rf directory
# ls la 做检查
