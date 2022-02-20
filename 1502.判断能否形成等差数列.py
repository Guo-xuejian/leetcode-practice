#
# @lc app=leetcode.cn id=1502 lang=python3
#
# [1502] 判断能否形成等差数列
#

# @lc code=start
class Solution:
    def canMakeArithmeticProgression(self, arr: List[int]) -> bool:
        length = len(arr)
        if length == 2:
            return True
        arr.sort()
        lap = arr[1] - arr[0]
        for i in range(2, length):
            if arr[i] - arr[i - 1] != lap:
                return False
        return True
# @lc code=end
