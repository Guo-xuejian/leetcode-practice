#
# @lc app=leetcode.cn id=728 lang=python3
#
# [728] 自除数
#

# @lc code=start
class Solution:
    def selfDividingNumbers(self, left: int, right: int) -> List[int]:
        def valid(number):
            for num in str(number):
                if num == "0" or number % int(num) != 0:
                    return False
            return True
        res = []
        for num in range(left, right + 1):
            if valid(num):
                res.append(num)
        return res
# @lc code=end

