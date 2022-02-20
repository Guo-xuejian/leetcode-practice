#
# @lc app=leetcode.cn id=1822 lang=python3
#
# [1822] 数组元素积的符号
#

# @lc code=start
class Solution:
    def arraySign(self, nums: List[int]) -> int:
        # 实际上只需要记录族中 zero_status 的正负情况
        positive = True
        zero_status = False
        for num in nums:
            if num == 0:
                zero_status = True
                break
            elif num < 0:
                if positive:
                    positive = False
                else:
                    positive = True
        if zero_status:
            return 0
        elif positive:
            return 1
        else:
            return -1
# @lc code=end

