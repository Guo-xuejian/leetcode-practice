#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
#

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # 遍历数组 双层遍历
        for i in range(len(nums)):
            for j in range(len(nums)):
                if i != j:
                    # 判定是否相加之和 == target
                    if nums[i] + nums[j] == target:
                        return [i, j]
# @lc code=end

