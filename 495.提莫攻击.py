#
# @lc app=leetcode.cn id=495 lang=python3
#
# [495] 提莫攻击
#

# @lc code=start
class Solution:
    def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
        # 存在重置的情况
        res = 0
        poinous_time = 0    # 当前中毒时间
        for i in range(len(timeSeries)):
            # 当前没有处于中毒状态
            if timeSeries[i] >= poinous_time:
                res += duration
            else:   # 中毒状态，结束时间减去前一次解毒时间
                res += timeSeries[i] + duration - poinous_time
            # 解毒时间重置
            poinous_time = timeSeries[i] + duration
        return res

# @lc code=end

