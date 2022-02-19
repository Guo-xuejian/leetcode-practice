#
# @lc app=leetcode.cn id=717 lang=python3
#
# [717] 1比特与2比特字符
#

# @lc code=start
class Solution:
    def isOneBitCharacter(self, bits: List[int]) -> bool:
        i, n = 0, len(bits)
        while i < n - 1:    # 最后一位固定不需要
            # 0 开头则为第一种
            # 1 开头则跳过下一位，0 结尾不需要担心越界
            i += bits[i] + 1
        # 大于则 False，等于则 True
        return i == n - 1
# @lc code=end
