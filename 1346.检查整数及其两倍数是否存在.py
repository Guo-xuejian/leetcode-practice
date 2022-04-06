#
# @lc app=leetcode.cn id=1346 lang=python3
#
# [1346] 检查整数及其两倍数是否存在
#

# @lc code=start
class Solution:
    def checkIfExist(self, arr: List[int]) -> bool:
        num_dict = {}
        for num in arr:
            if num not in num_dict:
                num_dict[num] = 1
            else:
                num_dict[num] += 1
            
            if num == 0 and num_dict[num] >= 2:
                return True
            elif num != 0 and (num * 2 in num_dict or num / 2 in num_dict):
                return True
        return False
# @lc code=end

