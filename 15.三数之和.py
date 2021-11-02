#
# @lc app=leetcode.cn id=15 lang=python3
#
# [15] 三数之和
#

# @lc code=start
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        numOfNumbers = len(nums)
        # 排序后方便处理
        nums.sort()
        result = list()

        # i ==> 第一个数字索引
        # j ==> 第二个数字索引
        # h ==> 第三个数字索引
        for i in range(numOfNumbers):
            # 此次遍历应该与上一次的数字不同，否则就是无效的
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            # h 指向最右端
            h = numOfNumbers - 1
            # 获得相反数用于后续计算
            target = -nums[i]

            # i+1 在 i 的右侧遍历，避免重复
            for j in range(i + 1, numOfNumbers):
                # 同样不能对重复值做判定（因为排过序，所以会快速排除相同数字）
                if j > i + 1 and nums[j] == nums[j - 1]:
                    continue
                # 持续缩小范围，结束条件为索引 j 和 h 对应的数字之和小于等于 target 或者索引出现重复，同时重复情况需要在下方做出处理
                while j < h and nums[j] + nums[h] > target:
                    h -= 1
                # 如果索引重复，则退出
                if j == h:
                    break
                # 不重复且两数之和满足条件，将该种情况写入anx
                if nums[j] + nums[h] == target:
                    result.append([nums[i], nums[j], nums[h]])
        return result
# @lc code=end

