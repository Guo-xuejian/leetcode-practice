#
# @lc app=leetcode.cn id=119 lang=python3
#
# [119] 杨辉三角 II
#

# @lc code=start
class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        pre_row = []
        # 维护一个前一行和当前行，节省一点空间
        for i in range(rowIndex + 1):
            current_row = []
            for j in range(i + 1):
                if j == 0 or j == i:
                    current_row.append(1)
                else:
                    current_row.append(pre_row[j]+ pre_row[j - 1])
            pre_row = current_row
        return pre_row
# @lc code=end

