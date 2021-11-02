#
# @lc app=leetcode.cn id=84 lang=python3
#
# [84] 柱状图中最大的矩形
#

# @lc code=start
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        length_of_heights = len(heights)
        if length_of_heights == 1:
            return heights[0]
        # 定义相邻的前后指针
        i = 0
        j = 1
        max_area = -1
        while j < length_of_heights:
            max_area = max(max_area, 2*min(heights[i], heights[j]))
            i += 1
            j += 1
        return max_area
# @lc code=end

