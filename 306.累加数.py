#
# @lc app=leetcode.cn id=306 lang=python3
#
# [306] 累加数
#

# @lc code=start
class Solution:
    def isAdditiveNumber(self, num: str) -> bool:
        def dfs(num, first, second):
            if not num:
                return True     # 空满足条件
            total = first + second
            length = len(str(total))
            if str(total) == num[:length]:
                return dfs(num[length:], second, total)
            return False

        for i in range(1, len(num) - 1):    # 遍历第一个数字的边界
            if num[0] == '0' and i > 1:
                break   # 首位可以为 0，但是 0 必须要单独一个数字, i > 1 而不是 > 0 是因为截取操作的右边界是不包含的，所以需要等于 1 才能拿到至少一个字符，后续就是大于了
            for j in range(i+1, len(num)):  # 遍历第二个数字的边界
                if j - 1 > 1 and num[i] == '0':
                    break   # 前导符不能为 0
                if dfs(num[j:], int(num[:i]), int(num[i:j])):
                    return True
        return False
# @lc code=end
