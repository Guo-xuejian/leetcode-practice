#
# @lc app=leetcode.cn id=1018 lang=python3
#
# [1018] 可被 5 整除的二进制前缀
#

# @lc code=start
class Solution:
    def prefixesDivBy5(self, nums: List[int]) -> List[bool]:
        res = [False for _ in nums]
        pre_num = 0
        for i in range(len(nums)):
            pre_num = ((pre_num << 1) + nums[i]) % 5
            # 由于 Python 理论上可以存储无限大的数字，还可以这样,主要是其他强类型语言会有 int 的最大值限制
            # pre_num = pre_num * 2
            # if pre_num % 5 == 0:
            #     res[i] = True
            if pre_num == 0:
                res[i] = True
        return res
# @lc code=end

