#
# @lc app=leetcode.cn id=539 lang=python3
#
# [539] 最小时间差
#

# @lc code=start
class Solution:
    def findMinDifference(self, timePoints: List[str]) -> int:
        timePoints = [int(x.split(":")[0])*60 + int(x.split(":")[1])
                      for x in timePoints]
        timePoints.sort()
        min_time = float('inf')
        length = len(timePoints)
        for i in range(1, length):  # 从首位开始向后比较
            curr = timePoints[i] - timePoints[i-1]
            if curr < min_time:
                min_time = curr
                if min_time == 0:
                    return 0
        # 相邻，也有可能跨天，将最晚时间和最早时间进行比较
        if timePoints[0] + 1440 - timePoints[-1] < min_time:
            min_time = timePoints[0] + 1440 - timePoints[-1]
        return min_time
# @lc code=end
