#
# @lc app=leetcode.cn id=952 lang=python3
#
# [952] 按公因数计算最大组件大小
#
# https://leetcode.cn/problems/largest-component-size-by-common-factor/description/
#
# algorithms
# Hard (37.60%)
# Likes:    78
# Dislikes: 0
# Total Accepted:    4.4K
# Total Submissions: 11.3K
# Testcase Example:  '[4,6,15,35]'
#
# 给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：
# 
# 
# 有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
# 只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
# 
# 
# 返回 图中最大连通组件的大小 。
# 
# 
# 
# 
# 
# 
# 示例 1：
# 
# 
# 
# 
# 输入：nums = [4,6,15,35]
# 输出：4
# 
# 
# 示例 2：
# 
# 
# 
# 
# 输入：nums = [20,50,9,63]
# 输出：2
# 
# 
# 示例 3：
# 
# 
# 
# 
# 输入：nums = [2,3,6,7,4,12,21,39]
# 输出：8
# 
# 
# 
# 
# 提示：
# 
# 
# 1 <= nums.length <= 2 * 10^4
# 1 <= nums[i] <= 10^5
# nums 中所有值都 不同
# 
# 
#

# @lc code=start
class UnionFind:
    def __init__(self, n: int):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def merge(self, x: int, y: int) -> None:
        x, y = self.find(x), self.find(y)
        if x == y:
            return
        if self.rank[x] > self.rank[y]:
            self.parent[y] = x
        elif self.rank[x] < self.rank[y]:
            self.parent[x] = y
        else:
            self.parent[y] = x
            self.rank[x] += 1

class Solution:
    def largestComponentSize(self, nums: List[int]) -> int:
        uf = UnionFind(max(nums) + 1)
        for num in nums:
            i = 2
            while i * i <= num:
                if num % i == 0:
                    uf.merge(num, i)
                    uf.merge(num, num // i)
                i += 1
        return max(Counter(uf.find(num) for num in nums).values())
# @lc code=end

