# 面试题 05.01. 插入
# 给定两个整型数字 N 与 M，以及表示比特位置的 i 与 j（i <= j，且从 0 位开始计算）。
# 编写一种方法，使 M 对应的二进制数字插入 N 对应的二进制数字的第 i ~ j 位区域，不足之处用 0 补齐。具体插入过程如图所示。

class Solution:
    def insertBits(self, N: int, M: int, i: int, j: int) -> int:
        # 做位运算，需要对应的几位都 0，这样做替换的时候就不用做判定
        #   将 n 的 i - j 位变为 0
        for k in range(i, j + 1):
            # ~ 取反，101 -> 010 ,后做位于运算，这样除了第k位置为0，其余为没有受到影响
            N &= ~(1 << k)
        # 将 m 左移 i 位 ,与 n 做或运算,等价于将 m 插入 i - j 位，后几位位或运算不会产生影响
        return N | (M << i)
