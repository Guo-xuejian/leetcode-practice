#
# @lc app=leetcode.cn id=377 lang=python3
#
# [377] 组合总和 Ⅳ
#

# @lc code=start
class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        # [1] 代表数字 0 只有一种组成方案就是不选
        dp = [1] + [0] * target
        # 对应索引代表该数字的组成方案
        for i in range(1, target + 1):
            # 计算出从 0---target 中每一个数字再当前nums限制下的组成方案数量
            for num in nums:
                if num <= i:    # 大于该索引（也就是该数字）的数字不会成为其组成元素
                    dp[i] += dp[i - num]
        return dp[target]
# @lc code=end
