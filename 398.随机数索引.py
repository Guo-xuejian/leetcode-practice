#
# @lc app=leetcode.cn id=398 lang=python3
#
# [398] 随机数索引
#

# @lc code=start
import random
class Solution:

    def __init__(self, nums: List[int]):
        self.nums = nums

    def pick(self, target: int) -> int:
        index_list = [x for x, v in enumerate(self.nums) if v == target]
        # 需要注意 random.randint 的两个参数是边界值，都是可达的
        return index_list[random.randint(0, len(index_list)-1)]


# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.pick(target)
# @lc code=end
