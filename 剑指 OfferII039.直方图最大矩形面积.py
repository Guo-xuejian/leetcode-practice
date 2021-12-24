# 给定非负整数数组 heights ，数组中的数字用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

# 求在该柱状图中，能够勾勒出来的矩形的最大面积。
# 输入：heights = [2,1,5,6,2,3]
# 输出：10
# 解释：最大的矩形为图中红色区域，面积为 10



class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        max_area = 0
        length = len(heights)
        for i in range(length):
            curr_height = heights[i]
            # 每一个高度都作为一个起点，遍历左右，找到两个单调栈的边界，计算面积与 max_area 比较后取较大值即可
            left = i - 1
            right = i + 1
            if length * curr_height > max_area:  # 这个判定可以帮助之后的遍历减少计算，由于之前的计算已经包含，或者说大于该直方面积，所以不需要再次进行额外计算，剪枝操作，有意思
                while left >= 0 and heights[left] >= curr_height:
                    left -= 1
                while right <= length -1 and heights[right] >= curr_height:
                    right += 1
                max_area = max(max_area, (right - left - 1) * curr_height)
        return max_area