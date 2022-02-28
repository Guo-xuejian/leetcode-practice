#
# @lc app=leetcode.cn id=1356 lang=python3
#
# [1356] 根据数字二进制下 1 的数目排序
#

# @lc code=start
class Solution:
    def sortByBits(self, arr: List[int]) -> List[int]:
        def countOnes(num):
            res = 0
            while num > 0:
                res += num & 1
                num >>= 1
            return res
        
        arr.sort(key = lambda ele: (countOnes(ele), ele))
        return arr
# @lc code=end

