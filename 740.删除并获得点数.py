#
# @lc app=leetcode.cn id=740 lang=python3
#
# [740] 删除并获得点数
#

# @lc code=start
class Solution:
    def deleteAndEarn(self, nums: List[int]) -> int:
        # 和打家劫舍一样，不能偷取相邻的,但是这里有一些区别，nums[i] - 1 和 nums[i] + 1 可能不存在
        # 那就找到最大值 maxNum 构建一个 [1,2,3,4,.....maxNum]
        if len(nums) == 1:
            return nums[0]

        def rob():  # 打家劫舍
            pre_num, curr_num = dp[0], max(dp[0], dp[1])
            for i in range(2, max_num + 1):
                pre_num, curr_num = curr_num, max(pre_num + dp[i], curr_num)
            return curr_num

        max_num = max(nums)
        dp = [0 for _ in range(max_num + 1)]
        for num in nums:
            dp[num] += num  # 重复无所谓，反正一起被偷

        return rob()
# @lc code=end
