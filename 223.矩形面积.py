#
# @lc app=leetcode.cn id=223 lang=python3
#
# [223] 矩形面积
#

# @lc code=start
class Solution:
    def computeArea(self, ax1: int, ay1: int, ax2: int, ay2: int, bx1: int, by1: int, bx2: int, by2: int) -> int:
        # 先计算出两个矩形的面积
        area1 = (ax2 - ax1) * (ay2 - ay1)
        area2 = (bx2 - bx1) * (by2 - by1)
        # 获取长宽，由于右顶点必定大于左顶点
        overlapWidth = min(ax2, bx2) - max(ax1, bx1)
        overlapHeight = min(ay2, by2) - max(ay1, by1)
        # 如果出现负数，则意味着没有共同区域， overlapArea 计算结果为 0
        overlapArea = max(overlapWidth, 0) * max(overlapHeight, 0)
        # 总面积减去覆盖面积
        return area1 + area2 - overlapArea
# @lc code=end

