#
# @lc app=leetcode.cn id=390 lang=python3
#
# [390] 消除游戏
#

# @lc code=start
class Solution:
    def lastRemaining(self, n: int) -> int:
        head = 1
        step = 1
        left = True
        
        while n > 1:
            # 从左边开始移除 or（从右边开始移除，数列总数为奇数）
            if left or n % 2 != 0:
                # 相当于只要是从左边开始删减，就等于给原有的第一个数字加上了 step
                # 12345 左删 ==> 24 加 2**1  24 右删 ===> 2 由于不是左所以头一个数字无影响
                # 同时又是每隔一个删除，其实单双无影响，反正总会在左右两种情况下被跳过！！！错
                # 还是有影响的，单数在右删情况下会删除第一个数字
                head += step
            
            step <<= 1 # 步长 * 2
            n >>= 1 # 总数 / 2
            left = not left #取反移除方向

        return head
# @lc code=end

