#
# @lc app=leetcode.cn id=384 lang=python3
#
# [384] 打乱数组
#

# @lc code=start
import random
class Solution:

    def __init__(self, nums: List[int]):
        self.nums = nums
        self.origin = nums.copy()

    def reset(self) -> List[int]:
        self.nums = self.origin
        return self.nums

    def shuffle(self) -> List[int]:
        # random.shuffle(self.nums)
        shuffled_list = [0 for _ in range(len(self.nums))]
        for i in range(len(shuffled_list)):
            j = random.randint(0,len(self.nums)-1)
            shuffled_list[i] = self.nums[j]
            self.nums = self.nums[:j] + self.nums[j+1:]
        self.nums = shuffled_list
        return self.nums


# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.reset()
# param_2 = obj.shuffle()
# @lc code=end
