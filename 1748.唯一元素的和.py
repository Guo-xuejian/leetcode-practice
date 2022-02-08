#
# @lc app=leetcode.cn id=1748 lang=python3
#
# [1748] 唯一元素的和
#

# @lc code=start
class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        return sum(num for num, cnt in Counter(nums).items() if cnt == 1)

# @lc code=end

