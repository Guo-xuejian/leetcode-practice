#
# @lc app=leetcode.cn id=118 lang=python3
#
# [118] 杨辉三角
#

# @lc code=start
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        result = list()
        for i in range(numRows):
            row = list()
            for j in range(0, i + 1):
                if j == 0 or j == i:    # 边界直接加入 1
                    row.append(1)
                else:   # 其余情况正常规则计算
                    row.append(result[i-1][j] + result[i-1][j-1])
            result.append(row)
        return result
# @lc code=end

