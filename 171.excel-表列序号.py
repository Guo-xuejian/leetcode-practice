#
# @lc app=leetcode.cn id=171 lang=python3
#
# [171] Excel 表列序号
#

# @lc code=start
class Solution:
    def titleToNumber(self, columnTitle: str) -> int:
        num, multiple = 0, 1
        length = len(columnTitle)
        for i in range(length - 1, -1, -1):
            k = ord(columnTitle[i]) - 64    # 计算该位数字
            num += k * multiple             # 乘上对应倍数
            multiple *= 26                  # 倍数随位数倍增 26
        return num
# @lc code=end

