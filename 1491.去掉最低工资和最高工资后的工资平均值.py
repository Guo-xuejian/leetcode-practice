#
# @lc app=leetcode.cn id=1491 lang=python3
#
# [1491] 去掉最低工资和最高工资后的工资平均值
#

# @lc code=start
class Solution:
    def average(self, salary: List[int]) -> float:
        max_num, min_num, total = salary[0], salary[0], 0
        for num in salary:
            if num > max_num:
                max_num = num
            elif num < min_num:
                min_num = num
            total += num
        return (total - max - min_num) / (len(salary) - 2)
# @lc code=end

