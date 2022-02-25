#
# @lc app=leetcode.cn id=1232 lang=python3
#
# [1232] 缀点成线
#

# @lc code=start
class Solution:
    def checkStraightLine(self, coordinates: List[List[int]]) -> bool:
        delta_x, delta_y = coordinates[0][0], coordinates[0][1]

        x1, y1 = coordinates[1][0] - delta_x, coordinates[1][1] - delta_y

        for point in coordinates[2:]:
            x2, y2 = point[0] - delta_x, point[1] - delta_y
            if x1 * y2 - x2 * y1 != 0:
                return False
        
        return True
# @lc code=end

