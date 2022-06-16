#
# @lc app=leetcode.cn id=532 lang=python3
#
# [532] 数组中的 k-diff 数对
#

# @lc code=start
class Solution:
    def findPairs(self, nums: List[int], k: int) -> int:
        visited, res = set(), set()
        for num in nums:
            if num - k in visited:
                res.add(num - k)
            if num + k in visited:
                res.add(num)
            visited.add(num)
        return len(res)
# @lc code=end

