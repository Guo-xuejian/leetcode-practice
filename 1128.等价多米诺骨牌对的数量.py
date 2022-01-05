#
# @lc app=leetcode.cn id=1128 lang=python3
#
# [1128] 等价多米诺骨牌对的数量
#

# @lc code=start
class Solution:
    def numEquivDominoPairs(self, dominoes: List[List[int]]) -> int:
        nums = [0] * 100
        res = 0
        for x, y in dominoes:
            val = (x*10 + y if x <= y else x + y*10)
            res += nums[val]    # 666 2==1 3==2 4==5 5==9,实际也就是每多出来一个，那么该元素与之前存在过的等价各自组合，也就是加上上一次的数量，6666 
            nums[val] += 1
        return res
# @lc code=end
