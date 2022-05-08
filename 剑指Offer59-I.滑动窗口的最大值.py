# 剑指 Offer 59 - I. 滑动窗口的最大值
# 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

# 示例:

# 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
# 输出: [3,3,5,5,6,7] 
# 解释: 

#   滑动窗口的位置                最大值
# ---------------               -----
# [1  3  -1] -3  5  3  6  7       3
#  1 [3  -1  -3] 5  3  6  7       3
#  1  3 [-1  -3  5] 3  6  7       5
#  1  3  -1 [-3  5  3] 6  7       5
#  1  3  -1  -3 [5  3  6] 7       6
#  1  3  -1  -3  5 [3  6  7]      7
 

# 提示：

# 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。



class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        n = len(nums)
        if n == 0:
            return []
        prefixMax, suffixMax = [0] * n, [0] * n
        for i in range(n):
            if i % k == 0:
                prefixMax[i] = nums[i]
            else:
                prefixMax[i] = max(prefixMax[i - 1], nums[i])
        for i in range(n - 1, -1, -1):
            if i == n - 1 or (i + 1) % k == 0:
                suffixMax[i] = nums[i]
            else:
                suffixMax[i] = max(suffixMax[i + 1], nums[i])

        ans = [max(suffixMax[i], prefixMax[i + k - 1]) for i in range(n - k + 1)]
        return ans