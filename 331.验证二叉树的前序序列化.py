#
# @lc app=leetcode.cn id=331 lang=python3
#
# [331] 验证二叉树的前序序列化
#

# @lc code=start
class Solution:
    def isValidSerialization(self, preorder: str) -> bool:
        # 将根节点视为某一个节点的左右子节点中的一个，每一个非空节点有两个空节点
        length = len(preorder)
        slots = 1
        index = 0
        while index < length:
            if slots == 0:  # 还在循环内就没有空节点空间的话就退出
                return False
            if preorder[index] == ",":  # 连接符
                index += 1
            elif preorder[index] == "#":    # 空节点占用
                index += 1
                slots -= 1
            else:
                # 非空节点出现，也占用空节点空间，但是由于非空，其本身也自带两个空节点的空间
                while index < length and preorder[index] != ",":    # 数字可能不止一位
                    index += 1
                slots += 1  # slots = slots - 1 + 2
        # 不允许多于的空节点出现
        return slots == 0
# @lc code=end
