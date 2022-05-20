#
# @lc app=leetcode.cn id=436 lang=python3
#
# [436] 寻找右区间
#

# @lc code=start
class Solution:
    def findRightInterval(self, intervals: List[List[int]]) -> List[int]:
        for i, interval in enumerate(intervals):
            interval.append(i)
        intervals.sort()

        n = len(intervals)
        ans = [-1] * n
        for _, end, id in intervals:
            i = bisect_left(intervals, [end])
            if i < n:
                ans[id] = intervals[i][2]
        return ans
# @lc code=end

