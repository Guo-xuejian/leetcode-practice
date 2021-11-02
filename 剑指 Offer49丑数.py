class Solution:
    def nthUglyNumber(self, n: int) -> int:
        # 维护一个列表，每个索引对应的值就是该顺序位置上的丑数，每一遍历n2,n3,n5中的最小值，如果多个都为最小值，那么都要加一，也就是该数字乘上对应的权重（2，3，5）等于当前的dp[i]，那么下一次再乘上权重就一定不会大于所需要的值
        dp = [1]*n
        a = 0
        b = 0
        c = 0
        for i in range(1, n):
            n2 = dp[a]*2
            n3 = dp[b]*3
            n5 = dp[c]*5
            dp[i] = min(n2, n3, n5)
            # 下方的代码不能使用elif，要做三次判断，因为但凡符合条件都要加一
            if n2 == dp[i]:
                a += 1
            if n3 == dp[i]:
                b += 1
            if n5 == dp[i]:
                c += 1
        return dp[-1]