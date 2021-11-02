class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        # 每一次遍历，将正数和该位置上的数字加总，负数则和0加总，也就维护了每个数字为该数字周围范围内的最大和，最终返回一个nums中的最大值即可
        for i in range(1, len(nums)):
            nums[i] += max(nums[i - 1], 0)
        return max(nums)