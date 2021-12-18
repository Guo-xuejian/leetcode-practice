#
# @lc app=leetcode.cn id=70 lang=python3
#
# [70] 爬楼梯
#

# @lc code=start
class Solution:
    def climbStairs(self, n: int) -> int:
        # if n <= 2:  # 小于等于 2 时，可以直接得出结果
        #     return n
        # # 其余情况递归得出总和==========但是会超出限制时间
        # return self.climbStairs(n-1) + self.climbStairs(n-2)

        # 动态规划，前者影响后者，类似于 Django 模板中的管道符 |
        # fib_list = [0 for _ in range(n+1)]
        # fib_list[0] = 1
        # fib_list[1] = 1
        # for i in range(2, n+1):
        #     fib_list[i] = fib_list[i-1] + fib_list[i-2]
        # return fib_list[-1]

        # 也可以不使用那么大的列表，直接使用变量交换
        pre = 1
        after = 1
        for _ in range(2, n+1):
            pre, after = after, after + pre
        return after
# @lc code=end

