class Solution:
    def translateNum(self, num: int) -> int:
        # 动态规划中的零位目前都是为了保证转移方程的完整性，实际上不对应数据，是为了完整而构建的数字，值根据转移方程推导而得
        str_num = str(num)
        a = 1
        b = 1
        for i in range(2, len(str_num) + 1):
            c = a + b if '10' <= str_num[i - 2: i] <= '25' else a
            b = a
            a = c
        return a