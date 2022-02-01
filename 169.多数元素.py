#
# @lc app=leetcode.cn id=169 lang=python3
#
# [169] 多数元素
#

# @lc code=start
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        vote, res = 0, 0    # res 当前优势方，多者为最终优势方
        for num in nums:
            if vote == 0:   # 同归于尽，出现新阵营时重置
                res = num
            vote += 1 if res == num else -1 # 同阵营 +1，否则 -1
        return res
# @lc code=end

