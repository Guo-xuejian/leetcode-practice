#
# @lc app=leetcode.cn id=1414 lang=python3
#
# [1414] 和为 K 的最少斐波那契数字数目
#

# @lc code=start
class Solution:
    def findMinFibonacciNumbers(self, k: int) -> int:
        fib_list = [1, 1]
        while fib_list[-1] < k:
            fib_list.append(fib_list[-1] + fib_list[-2])
        # 最少，肯定优先较大的值，那就斐波那契数列反着遍历，最大，次大，较小补充，因为存在 1 ，所以必定可以凑齐
        res, i = 0, len(fib_list) - 1
        while k:    # 为 0 退出
            if k >= fib_list[i]:    # 减去较大值，结果 +1
                k -= fib_list[i]
                res += 1
            i -= 1  # 指针前移，寻找下一个较大数字
        return res
# @lc code=end

