#
# @lc app=leetcode.cn id=1399 lang=python3
#
# [1399] 统计最大组的数目
#

# @lc code=start
import collections


class Solution:
    def countLargestGroup(self, n: int) -> int:
        hashMap = collections.Counter()
        for i in range(1, n + 1):
            key = sum([int(x) for x in str(i)])     # python 处理的确简单许多,代码也十分简洁
            hashMap[key] += 1
        maxValue = max(hashMap.values())    # 并列最多的组
        count = sum(1 for v in hashMap.values() if v == maxValue)   # 加总数量
        return count

# @lc code=end
