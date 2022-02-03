#
# @lc app=leetcode.cn id=1725 lang=python3
#
# [1725] 可以形成最大正方形的矩形数目
#

# @lc code=start
class Solution:
    def countGoodRectangles(self, rectangles: List[List[int]]) -> int:
        res, max_length = 0, 0
        for length, width in rectangles:
            num = min(length, width)    # 短边即为最大切分正方形边长
            if num == max_length:
                res += 1
            elif num > max_length:  # 更大时重置 res 和 max_kength
                res = 1
                max_length = num
        return res
# @lc code=end
