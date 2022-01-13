#
# @lc app=leetcode.cn id=74 lang=python3
#
# [74] 搜索二维矩阵
#

# @lc code=start
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        # 超出边界则 False
        if not (matrix[0][0] <= target or target <= matrix[-1][-1]):
            return False
        row = -1    # 两种情况，第一行是则row==0，最后一行则-1，一开始随便给的，还没想这么多，Go实现的时候不支持负数索引，报错了
        # 寻找包含该元素合适的区间
        if len(matrix) == 1:
            row = 0
        else:
            for i in range(len(matrix)-1):
                if matrix[i][0] <= target and target < matrix[i+1][0]:
                    row = i
                    break
        # 加上下方代码逻辑更完善，Python凑巧可以，其余不支持负数索引的语言加上
        # if matrix[-1][0] >= target:
        #     row = -1
        # 二分查找
        left, right = 0, len(matrix[row])-1
        while left <= right:
            mid = (left + right) // 2
            one = matrix[row][mid]
            if one == target:
                # 找到了！
                return True
            elif one < target:
                left = mid + 1
            else:
                right = mid - 1
        # 在合适范围却没有
        return False
# @lc code=end
