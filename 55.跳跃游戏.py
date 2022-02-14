#
# @lc app=leetcode.cn id=55 lang=python3
#
# [55] 跳跃游戏
#

# @lc code=start
class Solution:
    def canJump(self, nums: List[int]) -> bool:
        max_index = 0       # 初始化当前能到达最远的位置
        for i, jump in enumerate(nums):   # i 为当前位置，jump 是当前位置的跳数
            if max_index >= i and i + jump > max_index:  # 如果当前位置能到达，并且当前位置 + 跳数 > 最远位置
                max_index = i + jump  # 更新最远能到达位置
        return max_index >= len(nums) - 1
# @lc code=end
