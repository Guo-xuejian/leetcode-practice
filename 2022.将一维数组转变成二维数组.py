#
# @lc app=leetcode.cn id=2022 lang=python3
#
# [2022] 将一维数组转变成二维数组
#

# @lc code=start
class Solution:
    def construct2DArray(self, original: List[int], m: int, n: int) -> List[List[int]]:
        res = []
        if len(original) != m*n:
            return res
        index = 0
        for i in range(m):  # 外层写出内层列表
            temp = [0 for _ in range(n)]
            for j in range(n):  # 内层写出所有元素
                temp[j] = original[index]
                index += 1
            res.append(temp)
        return res

# @lc code=end

