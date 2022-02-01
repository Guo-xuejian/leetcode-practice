#
# @lc app=leetcode.cn id=240 lang=python3
#
# [240] 搜索二维矩阵 II
#

# @lc code=start
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        length1, length2 = len(matrix), len(matrix[0])
        for i in range(length1):
            if matrix[i][0] > target:   # 当前行最小值仍大于，同时大于上一行最大值，无法满足
                return False
            if matrix[i][length2-1] >= target:  # 在当前行，二分查找
                left, right = 0, length2-1
                while left <= right:
                    mid = int((left + right) / 2)
                    curr = matrix[i][mid]
                    if curr == target:
                        return True
                    elif curr > target:
                        right = mid - 1
                    else:
                        left = mid + 1
        # 最终未找到
        return False

# @lc code=end
