# 剑指 Offer 46. 把数字翻译成字符串
# 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

 

# 示例 1:

# 输入: 12258
# 输出: 5
# 解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
 

# 提示：

# 0 <= num < 231


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