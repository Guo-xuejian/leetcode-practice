#
# @lc app=leetcode.cn id=1588 lang=python3
#
# [1588] 所有奇数长度子数组的和
#

# @lc code=start
class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        res = 0
        n = len(arr)
        for start in range(n):
            length = 1
            while start + length < n + 1:   # +1 是为了最终能够取到最后一个元素
                for num in arr[start:start + length]:
                    res += num
                length += 2
        return res
# @lc code=end

