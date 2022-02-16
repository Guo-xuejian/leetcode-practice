#
# @lc app=leetcode.cn id=45 lang=python3
#
# [45] 跳跃游戏 II
#

# @lc code=start
class Solution:
    def jump(self, nums: List[int]) -> int:
        curr_position = len(nums) - 1
        res = 0
        while curr_position:
            for i in range(curr_position):
                if i + nums[i] >= curr_position:
                    curr_position = i
                    res += 1
                    break
        return res
# @lc code=end

