# 剑指 Offer II 072. 求平方根
# 给定一个非负整数 x ，计算并返回 x 的平方根，即实现 int sqrt(int x) 函数。

# 正数的平方根有两个，只输出其中的正数平方根。

# 如果平方根不是整数，输出只保留整数的部分，小数部分将被舍去。


# 示例 1:

# 输入: x = 4
# 输出: 2
# 示例 2:

# 输入: x = 8
# 输出: 2
# 解释: 8 的平方根是 2.82842...，由于小数部分将被舍去，所以返回 2


# 提示:

# 0 <= x <= 231 - 1

class Solution:
    def mySqrt(self, x: int) -> int:
        # 二分法
        if x <= 1:
            return x
        # 由于题目要求取整数部分，其实也就是找一个数 x 满足 h**2 < x 同时 (h+1)**2 > x
        left = 0
        right = x

        while left <= right:
            mid = (left+right)//2
            if mid**2 <= x and (mid+1)**2 > x:
                return mid
            elif mid**2 < x:  # 数值偏小，需要扩大
                left = mid + 1
            elif mid**2 > x:  # 数值偏大，需要缩小
                right = mid - 1
