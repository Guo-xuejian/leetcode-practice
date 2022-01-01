#
# @lc app=leetcode.cn id=2023 lang=python3
#
# [2023] 连接后等于目标字符串的字符串对
#

# @lc code=start
class Solution:
    def numOfPairs(self, nums: List[str], target: str) -> int:
        res = 0
        num_dict = {}
        for num in nums:
            if num not in num_dict:
                num_dict[num] = 1
            else:
                num_dict[num] += 1
        target_length = len(target)
        for i in range(1, target_length):
            p, s = target[:i], target[i:]   # 枚举所有前缀+后缀
            if p not in num_dict or s not in num_dict:
                continue
            if p != s:
                res += num_dict[p] * num_dict[s]
            else:
                # 前后缀相同时，相当于从 cnt[p] 个下标中选择两个不同下标的排列数，即 A(cnt[p], 2)
                res += num_dict[p] * (num_dict[p] - 1)
        return res
# @lc code=end
