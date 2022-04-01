#
# @lc app=leetcode.cn id=1385 lang=python3
#
# [1385] 两个数组间的距离值
#

# @lc code=start
class Solution:
    def findTheDistanceValue(self, arr1: List[int], arr2: List[int], d: int) -> int:
        def twoSideCheck(num):
            for i in range(0, d + 1):
                if num + i in num_dict or num - i in num_dict:
                    return False
            return True
        num_dict = {}
        for num in arr2:
            if num not in num_dict:
                num_dict[num] = 1
            else:
                num_dict[num] += 1
        res = 0
        for num in arr1:
            if twoSideCheck(num):
                res += 1
        return res
# @lc code=end

