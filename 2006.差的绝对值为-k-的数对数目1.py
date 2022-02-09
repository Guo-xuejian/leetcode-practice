#
# @lc app=leetcode.cn id=2006 lang=python3
#
# [2006] 差的绝对值为 K 的数对数目
#

# @lc code=start
class Solution:
    def countKDifference(self, nums: List[int], k: int) -> int:
        # 哈希表记录加速访问
        res = 0
        num_dict = {}
        for num in nums:
            if num - k in num_dict:
                res += num_dict[num - k]
            if num + k in num_dict:
                res += num_dict[num + k]
            
            if num in num_dict:
                num_dict[num] += 1
            else:
                num_dict[num] = 1
        return res
# @lc code=end

