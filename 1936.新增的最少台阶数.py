#
# @lc app=leetcode.cn id=1936 lang=python3
#
# [1936] 新增的最少台阶数
#

# @lc code=start
class Solution:
    def addRungs(self, rungs: List[int], dist: int) -> int:
        stairs_num= 0
        current_height = 0
        for height in rungs:
            stairs_num += (height - current_height - 1) // dist
            current_height = height
        return stairs_num
# @lc code=end