#
# @lc app=leetcode.cn id=690 lang=python3
#
# [690] 员工的重要性
#

# @lc code=start
"""
# Definition for Employee.
class Employee:
    def __init__(self, id: int, importance: int, subordinates: List[int]):
        self.id = id
        self.importance = importance
        self.subordinates = subordinates
"""

class Solution:
    def getImportance(self, employees: List['Employee'], id: int) -> int:
        employee_dict = {employee.id: employee for employee in employees}
        # 递归遍历
        def dfs(id):
            employee = employee_dict[id]
            total = employee.importance + sum(dfs(stuff_id) for stuff_id in employee.subordinates)
            return total
        
        return dfs(id)

        
# @lc code=end

