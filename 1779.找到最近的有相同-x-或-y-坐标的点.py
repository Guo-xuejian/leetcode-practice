#
# @lc app=leetcode.cn id=1779 lang=python3
#
# [1779] 找到最近的有相同 X 或 Y 坐标的点
#

# @lc code=start
class Solution:
    def nearestValidPoint(self, x: int, y: int, points: List[List[int]]) -> int:
        res = -1
        minDistance = float('inf')
        for i in range(len(points)):
            if points[i][0] == x or points[i][1] == y:
                currDistance = abs(points[i][0] - x) + abs(points[i][1] - y)
                if currDistance < minDistance:
                    minDistance = currDistance
                    res = i
        return res
# @lc code=end

