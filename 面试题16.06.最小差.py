# 给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差

# 示例：
# 输入：{1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
# 输出：3，即数值对(11, 8)


class Solution:
    def smallestDifference(self, a: List[int], b: List[int]) -> int:
        a.sort()
        b.sort()
        i, j, diff = 0, 0, 0
        res = float('inf')
        while i < len(a) and j < len(b):
            if a[i] < b[j]:
                diff = b[j] - a[i]
                i += 1
            else:
                diff = a[i] - b[j]
                j += 1
            if diff < res:
                res = diff
        return res
    
    def quick_sort(self, a):
        left, right = [], []
        first = a[0]
        for i in range(len(a)):
            if a[i] < first:
                left.append(a[i])
            else:
                right.append(a[i])
        return self.quick_sort(left) + [first] + self.quick_sort(right)
