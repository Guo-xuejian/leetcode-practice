#
# @lc app=leetcode.cn id=46 lang=python3
#
# [46] 全排列
#

# @lc code=start
class Solution:
    def __init__(self):
        self.index = 0
    
    def permute(self, nums: List[int]) -> List[List[int]]:
        length1 = len(nums)
        all_range_num = self.factorial(length1)
        res = [[0 for _ in range(length1)] for _ in range(all_range_num)]
        
        def backTrack(temp_nums, path, length):
            if length == 0:
                for i in range(length1):
                    res[self.index][i] = path[i]
                self.index += 1
                return
            for i in range(length):
                curr = temp_nums[i]
                path.append(curr)
                temp_nums = temp_nums[:i] + temp_nums[i+1:]
                backTrack(temp_nums, path, len(temp_nums))
                temp_nums = temp_nums[:i] + [curr] + temp_nums[i:]
                path.pop()
        
        backTrack(nums, [], length1)
        return res
    
    def factorial(self, num):
        res = 1
        while num != 0:
            res *= num
            num -= 1
        return res
        

# @lc code=end

