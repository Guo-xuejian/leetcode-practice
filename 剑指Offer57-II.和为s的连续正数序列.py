# 剑指 Offer 57 - II. 和为s的连续正数序列
# 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

# 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。


# 示例 1：

# 输入：target = 9
# 输出：[[2,3,4],[4,5]]
# 示例 2：

# 输入：target = 15
# 输出：[[1,2,3,4,5],[4,5,6],[7,8]]


# 限制：

# 1 <= target <= 10^5

class Solution:
    def findContinuousSequence(self, target: int) -> List[List[int]]:
        left, right, sum_now, res = 1, 2, 3, []
        while left < right:
            if sum_now == target:
                # 加入序列
                res.append(list(range(left, right + 1)))
            if sum_now >= target:   # 大于左边界右移，同时减去数组首位值
                sum_now -= left
                left += 1
            else:   # 小于右边界右移，扩展序列，加上扩张值
                right += 1
                sum_now += right
        return res
