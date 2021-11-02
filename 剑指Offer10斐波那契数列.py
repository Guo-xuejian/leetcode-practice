class Solution:
    def fib(self, n: int) -> int:
        # 常规思路
        # if n < 2:
        #     return n
        # return self.fib(n - 1) + self.fib(n - 2)

        # 动态规划
        if n == 0:
            return 0
        if n == 2 or n == 1:
            return 1
        first_list = [0 for _ in range(n)]
        first_list[0] = 1
        first_list[1] = 1
        for j in range(2, n):
            first_list[j] = (first_list[j - 1] + first_list[j - 2]) % 1000000007    # 题目要求对这个数字取模
        return first_list[n - 1]